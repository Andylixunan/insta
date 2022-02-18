package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/register", register)
	router.Run(":8080")
}

func register(c *gin.Context) {

}
