package service

import (
	"fmt"
	pb "github.com/jfeng45/grpcservice"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"
	"time"
)

const service_name_db_query_user = "db query user"

// CacheService struct
type CacheService struct {
	Storage map[string][]byte
}

// Get function
//func (c *CacheService) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
//	fmt.Println("start server side Get called: ")
//	//time.Sleep(3000*time.Millisecond)
//	key := req.GetKey()
//	value := c.Storage[key]
//	resp := &pb.GetResp{Value: value}
//	fmt.Println("Get called with return of value: ", value)
//	return resp, nil
//}

//Get function
func (cs *CacheService) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
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
	value := cs.Storage[key]
	fmt.Println("get called with return of value: ", value)
	resp := &pb.GetResp{Value: value}
	return resp, nil

}

func (cs *CacheService) Store(ctx context.Context, req *pb.StoreReq) (*pb.StoreResp, error) {
	key := req.Key
	value := req.Value
	if oldValue, ok := cs.Storage[key]; ok {
		cs.Storage[key] = value
		fmt.Printf(" key=%v already exist, old vale=%v|replaced with new value=%v\n", key, oldValue, cs.Storage)
	} else {
		cs.Storage[key] = value
		fmt.Printf(" key=%v not existing, add new value=%v\n", key, cs.Storage)
	}
	r := &pb.StoreResp{}
	return r, nil
}

func (cs *CacheService) Dump(*pb.DumpReq, pb.CacheService_DumpServer) error {
	return nil
}

