package gateway

import (
	"time"

	"github.com/Andylixunan/insta/global/config"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var configs *config.Config
var logger *zap.SugaredLogger

func NewServer(c *config.Config, l *zap.SugaredLogger) *gin.Engine {
	configs, logger = c, l
	router := gin.New()
	router.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))
	accountAPIGroup := router.Group("/account")
	{
		accountAPIGroup.POST("/register", register)
		accountAPIGroup.POST("/login", login)
	}
	return router
}
