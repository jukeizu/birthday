FROM golang:1.11 as build
WORKDIR /go/src/github.com/jukeizu/birthday
COPY Makefile go.mod go.sum ./
RUN make deps
ADD . .
RUN make build-linux
RUN echo "nobody:x:100:101:/" > passwd

FROM scratch
COPY --from=build /go/src/github.com/jukeizu/birthday/passwd /etc/passwd
COPY --from=build --chown=100:101 /go/src/github.com/jukeizu/birthday/bin/birthday .
USER nobody
ENTRYPOINT ["./birthday"]
