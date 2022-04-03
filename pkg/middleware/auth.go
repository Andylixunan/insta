package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	pb "github.com/Andylixunan/insta/api/proto/auth"
	"github.com/Andylixunan/insta/internal/auth"
	"github.com/Andylixunan/insta/pkg/config"
	"github.com/Andylixunan/insta/pkg/log"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	expectedScheme  = "bearer"
	headerAuthorize = "authorization"
)

func AuthMiddleware(config *config.Config, logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authClient, err := auth.NewClient(logger, config)
		if err != nil {
			logger.Infof("initialize auth client error: %v", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		tokenStr := c.GetString(headerAuthorize)
		if len(tokenStr) == 0 {
			logger.Infof("empty auth token for header: %v", headerAuthorize)
			c.AbortWithError(http.StatusForbidden, errors.New(fmt.Sprintf("empty auth token for header: %v", headerAuthorize)))
			return
		}
		splits := strings.SplitN(tokenStr, " ", 2)
		if len(splits) < 2 || !strings.EqualFold(splits[0], expectedScheme) {
			logger.Infof("Bad authorization string: %v", tokenStr)
			c.AbortWithError(http.StatusForbidden, errors.New(fmt.Sprintf("Bad authorization string: %v", tokenStr)))
			return
		}
		token := splits[1]
		resp, err := authClient.ValidateToken(c, &pb.ValidateTokenRequest{
			Token: token,
		})
		if err != nil {
			errStatus, _ := status.FromError(err)
			if errStatus.Code() == codes.Unauthenticated {
				logger.Infof("auth token invalid: %v", token)
				c.AbortWithError(http.StatusForbidden, errors.New(fmt.Sprintf("auth token invalid: %v", token)))
				return
			} else {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
		c.Set("UserID", resp.UserId)
		c.Next()
	}
}
