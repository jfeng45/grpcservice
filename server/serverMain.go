package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	pb "github.com/jfeng45/grpcservice"
	"github.com/jfeng45/grpcservice/server/middleware"
	"github.com/jfeng45/grpcservice/server/service"
	"github.com/opentracing/opentracing-go"
	openzipkin "github.com/openzipkin/zipkin-go-opentracing"
	"google.golang.org/grpc"
	"net"
)

const (
	endpoint_url = "http://localhost:9411/api/v1/spans"
	host_url = "localhost:5051"
	service_name_cache_server = "cache server"
	network = "tcp"
)

func newTracer () (opentracing.Tracer, error) {
	collector, err := openzipkin.NewHTTPCollector(endpoint_url)
	if err != nil {
		return nil, err
	}
	recorder :=openzipkin.NewRecorder(collector, true, host_url, service_name_cache_server)
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
	connection, err := net.Listen(network, host_url)
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
func initCache() pb.CacheServiceServer{
	s :=make(map[string][]byte)
	s["123"] = []byte{123}
	s["231"] = []byte{231}
	cs := service.CacheService{Storage: s}
	return middleware.BuildGetMiddleware(&cs)
}

