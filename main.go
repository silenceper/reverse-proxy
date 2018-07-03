package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"strings"
)

type reverseProxy struct {
}

func (rp *reverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	newHost := strings.Replace(r.Host, ".", "-", -1)
	newHost = strings.Replace(newHost, "_", "-", -1)
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	remote, err := url.Parse(scheme + newHost)
	if err != nil {
		panic(err)
	}
	log.Printf("proxy: %s TO %s", strings.Join([]string{scheme, r.Host, r.RequestURI}, ""), strings.Join([]string{scheme, newHost, r.RequestURI}, ""))
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	//被代理的服务器host和port
	err := http.ListenAndServe(":80", &reverseProxy{})
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	startServer()
}
