package main

import (
	"time"

	"github.com/Andylixunan/mini-instagram/global/config"
	"github.com/Andylixunan/mini-instagram/global/log"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var configs *config.Config
var logger *zap.SugaredLogger

func main() {
	var err error
	logger, err = log.New()
	defer logger.Sync()
	if err != nil {
		logger.Fatal(err)
	}
	configs, err = config.Load("../global/config.json")
	if err != nil {
		logger.Fatal(err)
	}
	router := gin.New()
	router.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))
	accountAPIGroup := router.Group("/account")
	{
		accountAPIGroup.POST("/register", register)
		accountAPIGroup.POST("/login", login)
	}
	err = router.Run(configs.Gateway.Port)
	if err != nil {
		logger.Fatal(err)
	}
}
