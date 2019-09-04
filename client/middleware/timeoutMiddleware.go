package middleware

import (
	pb "github.com/jfeng45/grpcservice"
	"golang.org/x/net/context"
	"log"
	"time"
)

const (
	get_timeout    = 200
	//get_timeout    = 5000
)

type TimeoutCallGet struct {
	Next callGetter
}

func (tcg *TimeoutCallGet) CallGet(ctx context.Context, key string, c pb.CacheServiceClient) ( []byte, error) {
	var cancelFunc context.CancelFunc
	var ch = make(chan bool)
	var err error
	var value []byte
	ctx, cancelFunc= context.WithTimeout(ctx, get_timeout*time.Millisecond)
	go func () {
		value, err = tcg.Next.CallGet(ctx, key, c)
		ch<- true
	} ()
	select {
		case <-ctx.Done():
			log.Println("ctx timeout")
			//ctx timeout, call cancelFunc to cancel all the sub-processes in the calling chain
			cancelFunc()
			err = ctx.Err()
		case <-ch:
			log.Println("call finished normally")
	}
	return value, err
}
