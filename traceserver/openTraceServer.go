package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	pb "github.com/jfeng45/grpcservice"
	"github.com/opentracing/opentracing-go"
	openzipkin "github.com/openzipkin/zipkin-go-opentracing"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"time"
)
const (
	endpointURL = "http://localhost:9411/api/v1/spans"
	hostUrl = "localhost:5051"
	service_name_cache_server = "cache server"
	service_name_db_query_user = "db query user"
	network = "tcp"
)

//CacheService struct
type CacheService struct {
	storage map[string][]byte
}

//Get function
func (c *CacheService) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	time.Sleep(5*time.Millisecond)
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		pctx := parent.Context()
		if tracer := opentracing.GlobalTracer(); tracer != nil {
			mysqlSpan := tracer.StartSpan(service_name_db_query_user, opentracing.ChildOf(pctx))
			defer mysqlSpan.Finish()
			//do some operations
			time.Sleep(time.Millisecond * 10)
		}
	}
	key := req.GetKey()
	value := c.storage[key]
	fmt.Println("get called with return of value: ", value)
	resp := &pb.GetResp{Value: value}
	return resp, nil

}

func (c *CacheService) Store(ctx context.Context, req *pb.StoreReq) (*pb.StoreResp, error) {
	key := req.Key
	value := req.Value

	if oldValue, ok := c.storage[key]; ok {
		c.storage[key] = value
		fmt.Printf(" key=%v already exist, old vale=%v|replaced with new value=%v\n", key, oldValue, c.storage)
	} else {
		c.storage[key] = value
		fmt.Printf(" key=%v not existing, add new value=%v\n", key, c.storage)
	}

	r := &pb.StoreResp{}
	return r, nil
}

func (c *CacheService) Dump(*pb.DumpReq, pb.CacheService_DumpServer) error {
	return nil

}
func newTracer () (opentracing.Tracer, error) {
	collector, err := openzipkin.NewHTTPCollector(endpointURL)
	if err != nil {
		return nil, err
	}
	recorder :=openzipkin.NewRecorder(collector, true, hostUrl, service_name_cache_server)
	tracer, err := openzipkin.NewTracer(
		recorder,
		openzipkin.ClientServerSameSpan(true))

	if err != nil {
		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)

	return tracer, nil
}

func main() {
	fmt.Println("starting server...")
	connection, err := net.Listen(network, hostUrl)
	if err != nil {
		panic(err)
	}
	tracer,err  := newTracer()
	if err != nil {
		panic(err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer,otgrpc.LogPayloads()),
		),
	}
	srv := grpc.NewServer(opts...)
	cs := initCache()
	pb.RegisterCacheServiceServer(srv, cs)

	err = srv.Serve(connection)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("server listening on port 5051")
	}

}
func initCache() *CacheService{
	s :=make(map[string][]byte)
	s["123"] = []byte{123}
	s["231"] = []byte{231}
	return &CacheService{storage: s}
}
