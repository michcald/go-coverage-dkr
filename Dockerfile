FROM golang:1.16.4-alpine3.12

COPY coverage.go /coverage.go

ADD entrypoint.sh /
RUN ["chmod", "+x", "/entrypoint.sh"]

CMD ["/entrypoint.sh"]
