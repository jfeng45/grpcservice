package middleware

import (
	pb "github.com/jfeng45/grpcservice"
	"golang.org/x/net/context"
	"log"
	"time"
)

const (
    retry_count    = 3
    retry_interval = 200
)

type RetryCallGet struct {
	Next callGetter
}

func (tcg *RetryCallGet) CallGet(ctx context.Context, key string, csc pb.CacheServiceClient) ( []byte, error) {
	var err error
	var value []byte
	for i:=0; i<retry_count; i++ {
		value, err = tcg.Next.CallGet(ctx, key, csc)
		log.Printf("Retry number %v|error=%v", i+1, err)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(retry_interval)*time.Millisecond)
	}
	return value, err
}
