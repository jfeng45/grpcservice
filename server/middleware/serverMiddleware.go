package middleware

import (
	pb "github.com/jfeng45/grpcservice"
	"golang.org/x/net/context"
)

type CacheServiceMiddleware struct {
	Next pb.CacheServiceServer
}

func BuildGetMiddleware(cs  pb.CacheServiceServer ) pb.CacheServiceServer {
	tm := ThrottleMiddleware{cs}
	csm := CacheServiceMiddleware{&tm}
	return &csm
}

func (csm *CacheServiceMiddleware) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	return csm.Next.Get(ctx, req)
}

func (csm *CacheServiceMiddleware) Store(ctx context.Context, req *pb.StoreReq) (*pb.StoreResp, error) {
	return csm.Next.Store(ctx, req)
}

func (csm *CacheServiceMiddleware) Dump(dr *pb.DumpReq, csds pb.CacheService_DumpServer) error {
	return csm.Next.Dump(dr, csds)
}