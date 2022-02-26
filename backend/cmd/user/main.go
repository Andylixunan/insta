package main

import (
	"net"

	pb "github.com/Andylixunan/insta/api/proto/user"
	"github.com/Andylixunan/insta/internal/user"
	"github.com/Andylixunan/insta/pkg/config"
	"github.com/Andylixunan/insta/pkg/dbcontext"
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
	if err != nil {
		logger.Fatal(err)
	}
	// connect to db
	connectionStr := dbcontext.GetDBConnectionStr(
		configs.User.DB.User,
		configs.User.DB.Passwd,
		configs.User.DB.DSN,
		configs.User.DB.Name,
		configs.User.DB.Options,
	)
	db, err := dbcontext.NewDBConnection(connectionStr)
	if err != nil {
		logger.Fatal(err)
	}
	repository := user.NewRepository(logger, db)
	// start the grpc server
	svr := grpc.NewServer()
	userServer := user.NewServer(logger, repository)
	pb.RegisterUserServiceServer(svr, userServer)
	reflection.Register(svr)
	lis, err := net.Listen("tcp", configs.User.Port)
	if err != nil {
		logger.Fatal(err)
	}
	svr.Serve(lis)
}
