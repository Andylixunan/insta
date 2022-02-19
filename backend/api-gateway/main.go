package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("../global/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.POST("/register", register)
	err = router.Run(viper.GetString("api-gateway.http-port"))
	if err != nil {
		log.Fatal(err)
	}
}

func register(c *gin.Context) {

}
