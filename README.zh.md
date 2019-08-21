## gRPC Service

其他语言：
### **[English](README.md)**

这是一个展示在Go微服务中使用OpenTracing和Zipkin实施分布式跟踪的项目。
我最初从[Alan Shreve的gRPC缓存服务](https://about.sourcegraph.com/go/grpc-in-production-alan-shreve)中获取代码，并在其中添加了分布式跟踪功能。

### 安装和运行

### 安装 Zipkin
https://zipkin.io/pages/quickstart

#### 安装程序

```
go get github.com/jfeng45/grpcservice
```

运行服务器
```
cd traceserver
go run openTraceServer.go
```
运行客户端
```
cd traceclient
go run openTraceclient.go
```
### 授权

[MIT](LICENSE.txt) 授权



