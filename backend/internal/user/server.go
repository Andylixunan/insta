package user

import (
	"context"

	pb "github.com/Andylixunan/insta/api/proto/user"
	"github.com/Andylixunan/insta/pkg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
	logger *log.Logger
	repo   Repository
}

func NewServer(logger *log.Logger, repo Repository) pb.UserServiceServer {
	return &server{
		logger: logger,
		repo:   repo,
	}
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplmented")
}
