package service

//import (
//	"fmt"
//	pb "github.com/jfeng45/grpcservice"
//	"github.com/opentracing/opentracing-go"
//	openzipkin "github.com/openzipkin/zipkin-go-opentracing"
//	"golang.org/x/net/context"
//	"log"
//	"time"
//)

import (
	pb "github.com/jfeng45/grpcservice"
	"golang.org/x/net/context"
)

const service_name_call_get = "callGet"

type CacheClient struct {

}

func (cc *CacheClient) CallGet(ctx context.Context, key string, csc pb.CacheServiceClient) ( []byte, error) {
	getReq:=&pb.GetReq{Key:key}
	getResp, err :=csc.Get(ctx, getReq )
	if err != nil {
		return nil, err
	}
	value := getResp.Value
	return value, err
}

func (cc *CacheClient) CallStore(key string, value []byte, client pb.CacheServiceClient) ( *pb.StoreResp, error) {
	storeReq := pb.StoreReq{Key: key, Value: value}
	storeResp, err := client.Store(context.Background(), &storeReq)
	if err != nil {
		return nil, err
	}
	return storeResp, err
}

//func (cc *CacheClient) CallGet(ctx context.Context, key string, csc pb.CacheServiceClient) ( []byte, error) {
//	span := opentracing.StartSpan(service_name_call_get)
//	sc := span.Context()
//	a :=fmt.Sprintf("context:%+v", span.Context())
//	log.Println("a:", a)
//	scZipkin :=sc.(openzipkin.SpanContext)
//	log.Printf("zipkinTraceId:%v", scZipkin.TraceID)
//	defer span.Finish()
//	time.Sleep(5*time.Millisecond)
//	// Put root span in context so it will be used in our calls to the client.
//	ctx = opentracing.ContextWithSpan(ctx, span)
//	//ctx := context.Background()
//	getReq :=&pb.GetReq{Key:key}
//	getResp, err :=csc.Get(ctx, getReq )
//	if err != nil {
//		return nil, err
//	}
//	value := getResp.Value
//	return value, err
//}

