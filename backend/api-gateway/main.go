package main

import (
	"log"

	"github.com/Andylixunan/mini-instagram/global/config"
	"github.com/gin-gonic/gin"
)

var configs *config.Config

func main() {
	var err error
	configs, err = config.Load("../global/config.json")
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	accountGroup := router.Group("/account")
	{
		accountGroup.POST("/register", register)
	}
	log.Println(configs)
	err = router.Run(configs.Gateway.Port)
	if err != nil {
		log.Fatal(err)
	}
}
