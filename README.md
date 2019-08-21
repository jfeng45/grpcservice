# gRPC Service

Other language: 
### **[中文](README.zh.md)**

This is the project to show how distributed tracing works with Go Microservice using OpenTracing and Zipkin. 
I originally took the code from [Alan Shreve's gRPC cache service](https://about.sourcegraph.com/go/grpc-in-production-alan-shreve), and added distributed tracing feature into it.

## Getting Started

### Installing Zipkin
https://zipkin.io/pages/quickstart

### Installing Application
```
go get github.com/jfeng45/grpcservice
```

Run Server
```
cd traceserver
go run openTraceServer.go
```
Run Client
```
cd traceclient
go run openTraceclient.go
```
## License

[MIT](LICENSE.txt) License



