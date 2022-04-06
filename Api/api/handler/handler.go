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

// func GetClaims(h *handlerV1, c *gin.Context) jwt.MapClaims {
// 	var (
// 		ErrUnauthorized = errors.New("unauthorized")
// 		claims          jwt.MapClaims
// 		err             error
// 	)

// 	Authorization.Token = c.GetHeader("Authorization")
// 	if c.Request.Header.Get("Authorization") == "" {
// 		c.JSON(http.StatusUnauthorized, models.ResponseError{
// 			Error: models.InternalServerError{
// 				Message: "Unauthorized request",
// 			},
// 		})
// 		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
// 		return nil
// 	}

// 	h.jwtHandler.Token = Authorization.Token
// 	claims, err = h.jwtHandler.ExtractClaims()
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, models.ResponseError{
// 			Error: models.InternalServerError{
// 				Message: "Unauthorized request",
// 			},
// 		})
// 		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
// 		return nil
// 	}
// 	return claims
// }
