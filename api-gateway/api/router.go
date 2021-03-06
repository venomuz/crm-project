package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/venomuz/crm-project/api-gateway/api/docs"
	v1 "github.com/venomuz/crm-project/api-gateway/api/handlers/v1"
	"github.com/venomuz/crm-project/api-gateway/config"
	"github.com/venomuz/crm-project/api-gateway/pkg/logger"
	"github.com/venomuz/crm-project/api-gateway/services"
	"github.com/venomuz/crm-project/api-gateway/storage/repo"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	RedisRepo      repo.RepositoryStorage
}

// New @BasePath /v1
// New ...
// @SecurityDefinitions.apikey BearerAuth
// @Description GetMyProfile
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Redis:          option.RedisRepo,
	})

	api := router.Group("/v1")
	api.POST("/users", handlerV1.CreateUser)
	api.GET("/users/:id", handlerV1.GetUser)
	//api.DELETE("/users/:id", handlerV1.DeleteUser)
	//api.POST("/users/check", handlerV1.CheckReg)
	//api.POST("/users/verify/:code", handlerV1.Verify)
	//api.GET("/users/login", handlerV1.Login)
	//api.GET("/users/get", handlerV1.GetUserWithToken)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
	return router
}
