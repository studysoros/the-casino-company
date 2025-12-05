package grpc_clients

import (
	"github.com/studysoros/the-casino-company/shared/env"
	pb "github.com/studysoros/the-casino-company/shared/proto/cashier"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type cashierServiceClient struct {
	Client pb.CashierServiceClient
	conn   *grpc.ClientConn
}

func NewCashierServiceClient() (*cashierServiceClient, error) {
	cashierServiceUrl := env.GetString("CASHIER_SERVICE_URL", "cashier-service:9092")

	// TODO: ADD tracing interceptors for observability.

	conn, err := grpc.NewClient(cashierServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewCashierServiceClient(conn)

	return &cashierServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (c *cashierServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
