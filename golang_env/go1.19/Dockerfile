#FROM golang:1.19

FROM alpine:latest as build

ENV GOLANG_VERSION=1.19
ENV PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ENV GOPATH=/go

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn


COPY --from=golang:1.19-alpine /usr/local/go/ /usr/local/go/

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache bash vim

WORKDIR /root
