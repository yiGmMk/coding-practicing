FROM golang:alpine as builder

WORKDIR /go/src
COPY server.go .

RUN go build -o httpd ./server.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/httpd .
RUN chmod +x /root/httpd

ENTRYPOINT ["/root/httpd"]