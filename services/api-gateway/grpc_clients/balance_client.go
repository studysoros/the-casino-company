package grpc_clients

import (
	"os"

	pb "github.com/studysoros/the-casino-company/shared/proto/balance"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type balanceServiceClient struct {
	Client pb.BalanceServiceClient
	conn   *grpc.ClientConn
}

func NewBalanceServiceClient() (*balanceServiceClient, error) {
	balanceServiceUrl := os.Getenv("BALANCE_SERVICE_URL")
	if balanceServiceUrl == "" {
		balanceServiceUrl = "balance-service:9093"
	}

	// TODO: ADD tracing interceptors for observability.

	conn, err := grpc.NewClient(balanceServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewBalanceServiceClient(conn)

	return &balanceServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (c *balanceServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
