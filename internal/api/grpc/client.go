package grpc

import (
	"github.com/hypertonyc/rpc-incrementor-service/internal/api/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IncrementClient struct {
	Service pb.IncrementServiceClient
	conn    *grpc.ClientConn
}

// Creates new gRPC client
func NewClient(address string) (*IncrementClient, error) {
	c := IncrementClient{}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c.conn = conn
	c.Service = pb.NewIncrementServiceClient(conn)

	return &c, nil
}

// Close client connection
func (c *IncrementClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
