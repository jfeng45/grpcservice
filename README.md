# gRPC Service

Other language: 
### **[中文](README.zh.md)**

This is the project to show how Go Microservice works, and it has two major functionalities. 
The first is how distributed tracing works with Go Microservice using OpenTracing and Zipkin. The second is service resilience, which includes the following:

* Timeout
* Retry
* Rate limitting
* Circuit Breaker
* Fault Injection
* Bulkhead

Please read the following articles for detail：

1. [How to Do Tracing in Go Microservice?](https://medium.com/swlh/distributed-tracing-for-go-microservice-with-opentracing-1fc1aec76b3e)

1. [Service Resilience in Go Microservice](https://medium.com/@jfeng45/service-resilience-in-go-microservice-52001e441020)

I originally took the code from [Alan Shreve's gRPC cache service](https://about.sourcegraph.com/go/grpc-in-production-alan-shreve), and added above features into it.

## Getting Started

### Installing Zipkin
Even without Zipkin, the rest of the application still works.

https://zipkin.io/pages/quickstart

### Installing Application
```
go get github.com/jfeng45/grpcservice
```

Run Server
```
cd server
go run serverMain.go
```
Run Client
```
cd client
go run clientMain.go
```
## License

[MIT](LICENSE.txt) License



