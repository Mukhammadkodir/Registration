package main

import (
	"github/RegisterTask/Api/api"
	"github/RegisterTask/Api/config"
	"github/RegisterTask/Api/pkg/logger"
	"github/RegisterTask/Api/services"
	"github/RegisterTask/Api/storage/repo"
)

func main() {
	var inMemStrg repo.InMemoryStorageI

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api-gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error")
	}

	server := api.New(api.Option{
		Conf:            cfg,
		Logger:          log,
		ServiceManager:  serviceManager,
		InMemoryStorage: inMemStrg,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server")
		panic(err)
	}

}
