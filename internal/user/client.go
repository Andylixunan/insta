package user

import (
	"context"
	"time"

	pb "github.com/Andylixunan/insta/api/proto/user"
	"github.com/Andylixunan/insta/pkg/config"
	"github.com/Andylixunan/insta/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(logger *log.Logger, conf *config.Config) (pb.UserServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, conf.User.Host+conf.User.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewUserServiceClient(conn), err
}
