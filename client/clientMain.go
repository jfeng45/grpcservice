package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	pb "github.com/jfeng45/grpcservice"
	"github.com/jfeng45/grpcservice/client/middleware"
	"github.com/jfeng45/grpcservice/client/service"
	"github.com/opentracing/opentracing-go"
	zipkintracer "github.com/openzipkin-contrib/zipkin-go-opentracing"
	openzipkin "github.com/openzipkin/zipkin-go-opentracing"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

const (
	endpoint_url = "http://localhost:9411/api/v1/spans"
	host_url = "localhost:5051"
	service_name_cache_client = "cache service client"
)

func newTracer () (opentracing.Tracer, zipkintracer.Collector, error) {
	collector, err := openzipkin.NewHTTPCollector(endpoint_url)
	if err != nil {
		return nil, nil, err
	}
	recorder :=openzipkin.NewRecorder(collector, true, host_url, service_name_cache_client)
	tracer, err := openzipkin.NewTracer(
		recorder,
		openzipkin.ClientServerSameSpan(true))

	if err != nil {
		return nil,nil,err
	}
	opentracing.SetGlobalTracer(tracer)

	return tracer,collector, nil
}

func main() {
	fmt.Println("client start")
	tracer, collector, err :=newTracer()
	if err != nil {
		panic(err)
	}
	defer collector.Close()
	connection, err := grpc.Dial(host_url,
		grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())),
	)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	client := pb.NewCacheServiceClient(connection)
	callServer(client)
	testCircuitBreaker(client)
}

func callServer(csc pb.CacheServiceClient) {
	ctx :=context.Background()
	key:="123"
	cc := service.CacheClient{}
	cg := middleware.BuildGetMiddleware(&cc)
	value, err := cg.CallGet(ctx, key, csc)
	if err != nil {
		fmt.Println("error call get:", err)
	} else {
		fmt.Printf("value=%v for key=%v\n", value, key)
	}
	key = "231"
	value, err = cg.CallGet(ctx, key, csc)
	if err != nil {
		fmt.Println("error call get:", err)
	} else {
		fmt.Printf("value=%v for key=%v\n", value, key)
	}
}

func testCircuitBreaker(csc pb.CacheServiceClient) {
	callServer(csc)
	time.Sleep(time.Duration(20*1000)*time.Millisecond)

	callServer(csc)
	time.Sleep(time.Duration(5*1000)*time.Millisecond)
	callServer(csc)
	time.Sleep(time.Duration(20*2000)*time.Millisecond)

	callServer(csc)
}
