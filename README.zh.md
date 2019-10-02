## gRPC Service

其他语言：
### **[English](README.md)**

这是一个展示Go微服务功能的项目。他主要有两大功能，第一是中如何使用OpenTracing和Zipkin实施分布式跟踪。第二是服务韧性（Service Resilience），它包括以下几点：
* 服务超时 （Timeout）
* 服务重试 （Retry）
* 服务限流（Rate Limitting）
* 熔断器 （Circuit Breaker）
* 故障注入（Fault Injection）
* 舱壁隔离技术（Bulkhead）

Please read the following articles for detail：

1. [Go全链路跟踪详解](https://zhuanlan.zhihu.com/p/79419529)

1. [Go微服务容错与韧性（Service Resilience）](https://zhuanlan.zhihu.com/p/81111394)

我最初从[Alan Shreve的gRPC缓存服务](https://about.sourcegraph.com/go/grpc-in-production-alan-shreve)中获取代码，并在其中添加了以上功能。

### 安装和运行

### 安装 Zipkin
即使没有安装Zipkin，程序的其他部分也能运行。

https://zipkin.io/pages/quickstart

#### 安装程序

```
go get github.com/jfeng45/grpcservice
```

运行服务器
```
cd server
go run serverMain.go
```
运行客户端
```
cd client
go run clientMain.go
```
### 授权

[MIT](LICENSE.txt) 授权



