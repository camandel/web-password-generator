FROM golang:1.11-alpine as build
LABEL maintainer=""

# Copy the local package files to the container's workspace.
ADD . /go/src/password-generator

# build & install server
WORKDIR /go/src/password-generator

RUN apk add git && \
    go get github.com/GeertJohan/go.rice/rice && \
    rice embed-go && \
    go build -o /go/bin/password-generator .

FROM alpine:latest
COPY --from=build /go/bin/password-generator /go/bin/password-generator

ENTRYPOINT ["/go/bin/password-generator"]
