package auth

import (
	pb "github.com/Andylixunan/insta/api/proto/auth"
	"github.com/Andylixunan/insta/pkg/log"
)

type server struct {
	pb.UnimplementedAuthServiceServer
	logger *log.Logger
}

func NewServer(logger *log.Logger) pb.AuthServiceServer {
	return &server{
		logger: logger,
	}
}
