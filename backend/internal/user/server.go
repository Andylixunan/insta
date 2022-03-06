package user

import (
	"context"

	pb "github.com/Andylixunan/insta/api/proto/user"
	"github.com/Andylixunan/insta/pkg/log"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	password, err := generateFromPassword(req.GetUser().GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to bcrypt generate password: %v", err)
	}
	user := &User{
		Username:        req.GetUser().GetUsername(),
		Password:        password,
		Nickname:        req.GetUser().GetNickname(),
		SelfDescription: req.GetUser().GetSelfDescription(),
		Avatar:          req.GetUser().GetAvatar(),
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	resp := &pb.CreateUserResponse{
		User: entityToProtobuf(user),
	}
	return resp, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id := req.GetId()
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	resp := &pb.GetUserResponse{
		User: entityToProtobuf(user),
	}
	return resp, nil
}

func (s *server) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.GetUserResponse, error) {
	username := req.GetUsername()
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user by username: %v", err)
	}
	resp := &pb.GetUserResponse{
		User: entityToProtobuf(user),
	}
	return resp, nil
}

func (s *server) GetUserByUsernameAndPassword(ctx context.Context, req *pb.GetUserByUsernameAndPasswordRequest) (*pb.GetUserResponse, error) {
	username := req.GetUsername()
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user by username: %v", err)
	}
	ok := isCorrectPassword(user.Password, req.GetPassword())
	if !ok {
		return nil, status.Errorf(codes.Internal, "incorrect password")
	}
	resp := &pb.GetUserResponse{
		User: entityToProtobuf(user),
	}
	return resp, nil
}

func generateFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func isCorrectPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func entityToProtobuf(user *User) *pb.User {
	return &pb.User{
		Id:              user.ID,
		Username:        user.Username,
		Password:        user.Password,
		Nickname:        user.Nickname,
		SelfDescription: user.SelfDescription,
		Avatar:          user.Avatar,
		CreatedAt:       timestamppb.New(user.CreatedAt),
		UpdatedAt:       timestamppb.New(user.UpdatedAt),
	}
}
