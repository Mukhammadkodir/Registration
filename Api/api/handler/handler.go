package handlers

import (
	"github/RegisterTask/Api/api/token"
	"github/RegisterTask/Api/config"
	"github/RegisterTask/Api/pkg/logger"
	"github/RegisterTask/Api/services"
	"github/RegisterTask/Api/storage/repo"
)

type handlerV1 struct {
	log             logger.Logger
	serviceManager  services.IServiceManager
	cfg             config.Config
	inMemoryStorage repo.InMemoryStorageI
	jwtHandler      token.JWTHandler
}

type HandlerV1Config struct {
	Logger          logger.Logger
	ServiceManager  services.IServiceManager
	Cfg             config.Config
	InMemoryStorage repo.InMemoryStorageI
	JwtHandler      token.JWTHandler
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:             c.Logger,
		serviceManager:  c.ServiceManager,
		cfg:             c.Cfg,
		inMemoryStorage: c.InMemoryStorage,
	}
}
