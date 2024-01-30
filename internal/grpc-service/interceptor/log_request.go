package interceptor

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
	"time"
)

func MiddlewareUnaryRequest(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	startTime := time.Now()

	data, err := handler(ctx, req)
	if err != nil {
		log.Errorf("[gRPC server] unary request: %v", err)
	}
	endTime := time.Now()
	fmt.Printf("%v:%v:%v       [gRPC server] | %v | unary request: %v",
		startTime.Hour(), startTime.Minute(), startTime.Second(),
		endTime.Sub(startTime),
		info.FullMethod)

	return data, err
}
