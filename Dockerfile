FROM            golang:1.16-alpine as builder
WORKDIR         /go/src/ultre.me/calcbiz
RUN             apk --no-cache --update add nodejs-npm make gcc g++ musl-dev openssl-dev git
RUN             go get -u github.com/gobuffalo/packr/packr
ENV             GO111MODULE=on
COPY            go.* /go/src/ultre.me/calcbiz/
RUN             go mod download
COPY            . /go/src/ultre.me/calcbiz/
RUN             make packr
RUN             make install

FROM            alpine:3.14
RUN             apk --no-cache --update add ca-certificates && update-ca-certificates
COPY            --from=builder /go/bin/calcbiz /bin/calcbiz
#COPY            ./static .
ENTRYPOINT      ["/bin/calcbiz"]
CMD             ["server"]
EXPOSE          9000 9001
