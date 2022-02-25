package main

import (
	"github.com/Andylixunan/mini-instagram/gateway"
	"github.com/Andylixunan/mini-instagram/global/config"
	"github.com/Andylixunan/mini-instagram/global/log"
)

func main() {
	var err error
	logger, err := log.New()
	defer logger.Sync()
	if err != nil {
		logger.Fatal(err)
	}
	configs, err := config.Load("../../configs/config.json")
	if err != nil {
		logger.Fatal(err)
	}
	router := gateway.NewServer(configs, logger)
	err = router.Run(configs.Gateway.Port)
	if err != nil {
		logger.Fatal(err)
	}
}
