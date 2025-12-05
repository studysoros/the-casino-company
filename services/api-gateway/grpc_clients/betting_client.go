package grpc_clients

import (
	"github.com/studysoros/the-casino-company/shared/env"
	pb "github.com/studysoros/the-casino-company/shared/proto/betting"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type bettingServiceClient struct {
	Client pb.BettingServiceClient
	conn   *grpc.ClientConn
}

func NewBettingServiceClient() (*bettingServiceClient, error) {
	bettingServiceUrl := env.GetString("BETTING_SERVICE_URL", "betting-service:9094")

	// TODO: ADD tracing interceptors for observability.

	conn, err := grpc.NewClient(bettingServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewBettingServiceClient(conn)

	return &bettingServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (c *bettingServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
