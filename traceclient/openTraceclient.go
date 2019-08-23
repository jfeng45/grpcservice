package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	pb "github.com/jfeng45/grpcservice"
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin-contrib/zipkin-go-opentracing"
	openzipkin "github.com/openzipkin/zipkin-go-opentracing"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)
const (
	endpointURL = "http://localhost:9411/api/v1/spans"
	hostUrl = "localhost:5051"
	service_name_cache_client = "cache service client"
	service_name_call_get = "callGet"
)

func callStore(key string, value []byte, client pb.CacheServiceClient) ( *pb.StoreResp, error) {
	storeReq := pb.StoreReq{Key: key, Value: value}
	storeResp, err := client.Store(context.Background(), &storeReq)
	return storeResp, err
}

func callGet(key string, c pb.CacheServiceClient) ( []byte, error) {
	span := opentracing.StartSpan(service_name_call_get)
	sc := span.Context()
	a :=fmt.Sprintf("context:%+v", span.Context())
	log.Println("a:", a)
	scZipkin :=sc.(openzipkin.SpanContext)
	log.Printf("zipkinTraceId:", scZipkin.TraceID)
	defer span.Finish()
	time.Sleep(5*time.Millisecond)
	// Put root span in context so it will be used in our calls to the client.
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	//ctx := context.Background()
	getReq:=&pb.GetReq{Key:key}
	getResp, err :=c.Get(ctx, getReq )
	value := getResp.Value
	return value, err
}

func newTracer () (opentracing.Tracer, zipkintracer.Collector, error) {
	collector, err := openzipkin.NewHTTPCollector(endpointURL)
	if err != nil {
		return nil, nil, err
	}
	recorder :=openzipkin.NewRecorder(collector, true, hostUrl, service_name_cache_client)
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
	key:="123"
	tracer, collector, err :=newTracer()
	if err != nil {
		panic(err)
	}
	defer collector.Close()
	connection, err := grpc.Dial(hostUrl,
		grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())),
		)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	client := pb.NewCacheServiceClient(connection)
	value, err := callGet(key, client)
	if err != nil {
		fmt.Println("error call get:", err)
	} else {
		fmt.Println("value1:", value)
	}
	key = "231"
	value, err = callGet(key, client)
	if err != nil {
		fmt.Println("error call get:", err)
	} else {
		fmt.Println("value2:", value)
	}
}
