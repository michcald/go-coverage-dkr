FROM golang:1.15-alpine3.11

COPY coverage.go /coverage.go

ADD entrypoint.sh /
RUN ["chmod", "+x", "/entrypoint.sh"]

CMD ["/entrypoint.sh"]
