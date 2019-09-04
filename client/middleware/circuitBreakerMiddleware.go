package middleware

import (
	pb "github.com/jfeng45/grpcservice"
	"github.com/sony/gobreaker"
	"golang.org/x/net/context"
	"log"
)
var cb *gobreaker.CircuitBreaker

type CircuitBreakerCallGet struct {
	Next callGetter
}
func init() {
	var st gobreaker.Settings
	st.Name = "CircuitBreakerCallGet"
	st.MaxRequests = 2
	st.Timeout = 10
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 2 && failureRatio >= 0.6
	}
	cb = gobreaker.NewCircuitBreaker(st)
}

func (tcg *CircuitBreakerCallGet) CallGet(ctx context.Context, key string, c pb.CacheServiceClient) ( []byte, error) {
	var err error
	var value []byte
	var serviceUp bool
	log.Printf("state:%v", cb.State().String())
	_, err = cb.Execute(func() (interface{}, error) {
		value, err = tcg.Next.CallGet(ctx, key, c)
		if err != nil {
			return nil, err
		}
		serviceUp = true
		return value, nil
	})
	if !serviceUp {
		//return a default value here. You can also run a downgrade function here
		log.Printf("circuit breaker return error:%v\n", err)
		return []byte{0}, nil
	}
	return value, nil
}
