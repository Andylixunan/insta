package main

import (
	"github.com/Andylixunan/insta/gateway"
	"github.com/Andylixunan/insta/pkg/config"
	"github.com/Andylixunan/insta/pkg/log"
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
