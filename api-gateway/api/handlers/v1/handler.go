package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/venomuz/crm-project/api-gateway/api/model"
	Ijwt "github.com/venomuz/crm-project/api-gateway/api/token"
	"github.com/venomuz/crm-project/api-gateway/config"
	"github.com/venomuz/crm-project/api-gateway/pkg/logger"
	"github.com/venomuz/crm-project/api-gateway/services"
	"github.com/venomuz/crm-project/api-gateway/storage/repo"
	"net/http"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	redisStorage   repo.RepositoryStorage
	jwtHandler     Ijwt.JwtHendler
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	Redis          repo.RepositoryStorage
	jwtHandler     Ijwt.JwtHendler
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		redisStorage:   c.Redis,
		jwtHandler:     c.jwtHandler,
	}
}
func CheckClaims(h *handlerV1, c *gin.Context) (jwt.MapClaims, error) {
	var (
		ErrUnauthorized = errors.New("unauthorized")
		authorization   model.JwtReqMod
		claims          jwt.MapClaims
		err             error
	)

	authorization.Token = c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, ErrUnauthorized)
		h.log.Error("Unauthorized request:", logger.Error(err))

	}
	h.jwtHandler.Token = authorization.Token
	claims, err = h.jwtHandler.ExtractClaims()
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrUnauthorized)
		h.log.Error("token is invalid:", logger.Error(err))
		return nil, err
	}
	return claims, nil
}
