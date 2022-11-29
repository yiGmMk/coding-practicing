# builder模式

- 第一阶段，构建负责编译源码的构建者镜像
- 第二阶段，将第一阶段的输出作为输入，构建出最终的目标镜像。

## 多阶段构建

使用builder模式,通过as builder指定构建镜像,在构建完成后将产物拷贝到需要的镜像以减小镜像体积

```dockerfile
FROM golang:alpine as builder

WORKDIR /go/src
COPY server.go .

RUN go build -o httpd ./server.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/httpd .
RUN chmod +x /root/httpd

ENTRYPOINT ["/root/httpd"]
```
