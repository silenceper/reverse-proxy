package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
)

type reverseProxy struct {
}

func (rp *reverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	remote, err := url.Parse(scheme + r.Host)
	if err != nil {
		panic(err)
	}
	log.Printf("proxy: %s ", r.RequestURI)
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
