package main

import (
	"net/http"

	"github.com/Andylixunan/mini-instagram/global/client"
	"github.com/Andylixunan/mini-instagram/global/proto/account"
	"github.com/gin-gonic/gin"
)

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func register(c *gin.Context) {
	var registerInfo Register
	if err := c.ShouldBindJSON(&registerInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	grpcClient, err := client.NewAccountClient(configs.Account.Port)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := grpcClient.Register(c, &account.RegisterRequest{
		User: &account.User{
			Username: registerInfo.Username,
			Password: registerInfo.Password,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"userID": resp.GetUserID()})
}

func login(c *gin.Context) {

}
