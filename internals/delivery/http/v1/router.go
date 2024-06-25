package v1

import (
	"github.com/go-frame/internals/entity/proto"
	"github.com/go-frame/internals/lib"
	"github.com/go-frame/internals/lib/logger"
	"github.com/go-frame/internals/repository"
	"github.com/go-frame/internals/service/auth"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-frame/config"
	"github.com/go-frame/internals/entity"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

// Setup all routers
func SetupRouters(c *echo.Echo, conf *config.Config, db *bun.DB, jwtConfig middleware.JWTConfig, logger logger.Logger, sess *session.Session, usersServiceConn *grpc.ClientConn, redisClient *redis.Client) {

	usersServiceClient := proto.NewUserServiceClient(usersServiceConn)
	usersServiceRepo := repository.NewUsersServiceRepo(usersServiceClient)
	loginAttemptsRepo := repository.NewLoginAttemptRepo(db)
	authService := auth.NewService(loginAttemptsRepo, logger, conf, usersServiceRepo, redisClient)
	authHandler := NewAuthHandler(authService, logger)

	v1 := c.Group("/api/v1/auth")

	health := v1.Group("/health")

	authGroup := v1.Group("")

	//companyGroup := v1.Group("")

	//restricted := c.Group("", middleware.JWTWithConfig(jwtConfig))
	//authenticated := middleware.JWTWithConfig(jwtConfig)

	health.GET("", func(c echo.Context) error {
		logger.Infof("Health check RequestID: %s", lib.GetRequestID(c))
		return c.JSON(http.StatusOK, &entity.Response{
			Success: true,
			Message: "go erp core service Server is running properly",
			Status:  http.StatusOK,
		})
	})

	authHandler.MapAuthRoutes(authGroup)

}
