FROM golang:1.11-alpine as build

ADD . /go/src/web-password-generator

WORKDIR /go/src/web-password-generator

RUN apk add git && \
    go get github.com/GeertJohan/go.rice/rice && \
    rice embed-go && \
    go build -o /go/bin/web-password-generator .

FROM alpine:3.8
COPY --from=build /go/bin/web-password-generator /go/bin/web-password-generator

EXPOSE 8080
ENTRYPOINT ["/go/bin/web-password-generator"]
