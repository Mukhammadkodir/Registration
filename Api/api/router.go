package api

import (
	v1 "github/RegisterTask/Api/api/handler"
	"github/RegisterTask/Api/config"
	"github/RegisterTask/Api/pkg/logger"
	"github/RegisterTask/Api/services"
	"github/RegisterTask/Api/storage/repo"

	"github/RegisterTask/Api/api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @termsOfService	http://swagger.io/terms/

// @securityDefinitions.apikey BearerAuth
// @in  header
// @name Authorization

// @contact.name	Api Support
// @contact.url		http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/license/LICENSE-2.0.html

// @host	localhost:8080
// @BasePath /v1

type Option struct {
	Conf            config.Config
	Logger          logger.Logger
	ServiceManager  services.IServiceManager
	InMemoryStorage repo.InMemoryStorageI
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/v1"

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:          option.Logger,
		ServiceManager:  option.ServiceManager,
		Cfg:             option.Conf,
		InMemoryStorage: option.InMemoryStorage,
	})

	api := router.Group("/v1")

	api.POST("/reg", handlerV1.Register)
	api.GET("/user/:id", handlerV1.Get)
	api.POST("/login", handlerV1.Login)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
