package main

import (
	"context"
	"log"
	"net"

	"github.com/Andylixunan/mini-instagram/global/config"
	pb "github.com/Andylixunan/mini-instagram/global/proto/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var configs *config.Config

type server struct {
	pb.UnimplementedAccountServer
}

func (*server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (*server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func main() {
	var err error
	configs, err = config.Load("../../global/config.json")
	if err != nil {
		log.Fatal(err)
	}
	svr := grpc.NewServer()
	pb.RegisterAccountServer(svr, &server{})
	reflection.Register(svr)
	lis, err := net.Listen("tcp", configs.Account.Port)
	if err != nil {
		log.Fatal(err)
	}
	svr.Serve(lis)
}
