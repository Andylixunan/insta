package auth

import (
	"context"

	pb "github.com/Andylixunan/insta/api/proto/auth"
	"github.com/Andylixunan/insta/pkg/jwt"
	"github.com/Andylixunan/insta/pkg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewServer(logger *log.Logger, jwtManager *jwt.Manager) pb.AuthServiceServer {
	return &server{
		logger:     logger,
		jwtManager: jwtManager,
	}
}

type server struct {
	pb.UnimplementedAuthServiceServer
	logger     *log.Logger
	jwtManager *jwt.Manager
}

func (s *server) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	token, err := s.jwtManager.Generate(req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}
	return &pb.GenerateTokenResponse{
		Token: token,
	}, nil
}

func (s *server) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	claims, err := s.jwtManager.Validate(req.GetToken())
	if err != nil || claims.ID == 0 {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	return &pb.ValidateTokenResponse{
		Valid: true,
	}, nil
}
