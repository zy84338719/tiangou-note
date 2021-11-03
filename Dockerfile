FROM golang:alpine
WORKDIR /go/src/server
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o server .

FROM alpine:latest
LABEL MAINTAINER="zhangyi@murphyyi.com"

WORKDIR /go/src/server

COPY --from=0 /go/src/server ./

EXPOSE 59998

ENTRYPOINT ./server