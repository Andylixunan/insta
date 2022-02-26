package gateway

import (
	"github.com/Andylixunan/insta/pkg/log"

	"github.com/Andylixunan/insta/pkg/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	config *config.Config
	logger *log.Logger
}

func NewServer(config *config.Config, logger *log.Logger) *Server {
	server := &Server{
		Engine: gin.Default(),
		config: config,
		logger: logger,
	}
	accountAPIGroup := server.Group("/account")
	{
		accountAPIGroup.POST("/register", register)
		accountAPIGroup.POST("/login", login)
	}
	return server
}
