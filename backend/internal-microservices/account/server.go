package main

import (
	"context"
	"net"

	"github.com/Andylixunan/insta/global/config"
	dbcon "github.com/Andylixunan/insta/global/db"
	"github.com/Andylixunan/insta/global/log"
	"github.com/Andylixunan/insta/global/model"
	pb "github.com/Andylixunan/insta/global/proto/account"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var configs *config.Config
var db *gorm.DB
var logger *zap.SugaredLogger

type server struct {
	pb.UnimplementedAccountServer
}

func (*server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	username := req.GetUser().GetUsername()
	user := &model.User{}
	db.Where("username = ?", username).Find(user)
	if user.ID != 0 {
		logger.Infof("username already exists: %v, id: %v", user.Username, user.ID)
		return nil, status.Errorf(codes.AlreadyExists, "username already exists")
	}
	password, err := generateFromPassword(req.GetUser().GetPassword())
	if err != nil {
		logger.Warnf("failed to generate password hash, error: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to generate password hash")
	}
	insertedUser := &model.User{
		Username: username,
		Password: password,
	}
	db.Create(insertedUser)
	logger.Infof("user created: %v, ID: %v", insertedUser.Username, insertedUser.ID)
	return &pb.RegisterResponse{
		UserID: insertedUser.ID,
	}, nil
}

func (*server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func generateFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// func isCorrectPassword(hash, password string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

func main() {
	var err error
	// load logger
	logger, err = log.New()
	if err != nil {
		logger.Fatal(err)
	}
	defer logger.Sync()
	// load config
	configs, err = config.Load("../../global/config.json")
	if err != nil {
		logger.Fatal(err)
	}
	// connect to db
	connectionStr := dbcon.GetDBConnectionStr(
		configs.Account.DB.User,
		configs.Account.DB.Passwd,
		configs.Account.DB.DSN,
		configs.Account.DB.Name,
		configs.Account.DB.Options,
	)
	db, err = dbcon.NewDBConnection(connectionStr)
	if err != nil {
		logger.Fatal(err)
	}
	// start the grpc server
	svr := grpc.NewServer()
	pb.RegisterAccountServer(svr, &server{})
	reflection.Register(svr)
	lis, err := net.Listen("tcp", configs.Account.Port)
	if err != nil {
		logger.Fatal(err)
	}
	svr.Serve(lis)
}
