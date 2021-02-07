FROM golang:1.13-alpine as builder

RUN apk add --no-cache git openssl

ENV GOPATH=/go PATH=/go/bin:$PATH CGO_ENABLED=0 GO111MODULE=on
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

COPY . /src
WORKDIR /src

RUN go build main.go
RUN chmod +x main


FROM alpine:latest as runner
RUN apk add --no-cache ca-certificates

COPY --from=builder /src/main /usr/local/bin/main

RUN adduser -D iamgroot
USER iamgroot

ENTRYPOINT [ "/usr/local/bin/main" ]