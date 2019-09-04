package middleware

import (
	pb "github.com/jfeng45/grpcservice"
	"golang.org/x/net/context"
)

type callGetter interface {
	CallGet(ctx context.Context, key string, c pb.CacheServiceClient) ( []byte, error)
}
type CallGetMiddleware struct {
	Next callGetter
}
func BuildGetMiddleware(cc callGetter) callGetter {
	cbcg := CircuitBreakerCallGet{cc}
	tcg := TimeoutCallGet{&cbcg}
	rcg := RetryCallGet{&tcg}
	return &rcg
}

func (cg *CallGetMiddleware) CallGet(ctx context.Context, key string, csc pb.CacheServiceClient) ( []byte, error) {
	return cg.Next.CallGet(ctx, key, csc)
}

//func BuildGetMiddleware() callGetter {
//	cc := service.CacheClient{}
//	//tcg := TimeoutCallGet{&cc}
//	//rcg := RetryCallGet{&tcg}
//	return &cc
//}



