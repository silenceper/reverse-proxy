FROM golang:alpine as builder
ADD . /go/src/github.com/silenceper/reverse-proxy/
RUN cd /go/src/github.com/silenceper/reverse-proxy/ \
  && go get -v \
  && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine
MAINTAINER silenceper <silenceper@gmail.com>
COPY --from=builder /go/src/github.com/silenceper/reverse-proxy/app /bin/app
ENTRYPOINT ["/bin/app"]
EXPOSE 80
