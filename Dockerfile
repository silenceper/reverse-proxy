FROM alpine

ADD app /bin/

ENTRYPOINT ["app"]

EXPOSE 80
