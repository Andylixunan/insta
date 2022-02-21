package client

import (
	"context"
	"time"

	accountpb "github.com/Andylixunan/mini-instagram/global/proto/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAccountClient(port string) (accountpb.AccountClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := accountpb.NewAccountClient(conn)
	return client, nil
}
