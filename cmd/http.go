package cmd

import (
	v1 "github.com/go-frame/internals/delivery/http/v1"
	middleware2 "github.com/go-frame/internals/delivery/middleware"
	"github.com/go-frame/internals/lib/logger"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-frame/config"
	"github.com/go-frame/internals/entity"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/uptrace/bun"
)

func setupHTTPServer(conf *config.Config, db *bun.DB, sess *session.Session, appLogger logger.Logger, redisClient *redis.Client) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	if conf.Debug {
		e.Use(middleware.Logger())
	}

	jwtConfig := middleware.JWTConfig{
		Claims:       &entity.JwtClaim{},
		SigningKey:   []byte(conf.JwtSecret),
		ErrorHandler: middleware2.InvalidJwt,
	}

	usersServiceConn, err := grpc.Dial(conf.UsersServiceGrpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	v1.SetupRouters(e, conf, db, jwtConfig, appLogger, sess, usersServiceConn, redisClient)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}

func startHTTPServer(e *echo.Echo, httpConfig config.HTTP) {
	if err := e.Start(httpConfig.HTTPAddress); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server")
	}
}
