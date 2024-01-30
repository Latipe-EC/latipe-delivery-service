package grpcclient

import (
	"delivery-service/config"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	conn *grpc.ClientConn
}

func NewGrpcServerConnection(cfg *config.Config) *GrpcClient {
	conn, err := grpc.Dial(cfg.GRPC.Connection, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil
	}

	return &GrpcClient{conn: conn}
}

func (c GrpcClient) GetClient() *grpc.ClientConn {
	return c.conn
}

func (c GrpcClient) Close() error {
	return c.conn.Close()
}
