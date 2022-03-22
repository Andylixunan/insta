package main

import (
	"net"

	pb "github.com/Andylixunan/insta/api/proto/auth"
	"github.com/Andylixunan/insta/internal/auth"
	"github.com/Andylixunan/insta/pkg/config"
	"github.com/Andylixunan/insta/pkg/jwt"
	"github.com/Andylixunan/insta/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var err error
	// load logger
	logger, err := log.New()
	if err != nil {
		logger.Fatal(err)
	}
	defer logger.Sync()
	// load config
	configs, err := config.Load("../../configs/config.json")
	manager := jwt.NewManager(configs)
	// start the grpc server
	authServiceServer := auth.NewServer(logger, manager)
	if err != nil {
		logger.Fatal(err)
	}
	svr := grpc.NewServer()
	pb.RegisterAuthServiceServer(svr, authServiceServer)
	reflection.Register(svr)
	logger.Infof("auth server starts listening at %v", configs.Auth.Host+configs.Auth.Port)
	lis, err := net.Listen("tcp", configs.Auth.Host+configs.Auth.Port)
	if err != nil {
		logger.Fatal(err)
	}
	svr.Serve(lis)
}
