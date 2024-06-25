package cmd

import (
	"context"
	"fmt"
	"github.com/go-frame/config"
	"github.com/go-frame/config/database"
	"github.com/go-frame/internals/lib/logger"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func readConfigFileFromEnv() *config.Config {
	conf := config.NewConfig("./config.env")
	if conf == nil {
		panic("Failed to load configuration")
	}
	fmt.Println("Configuration loaded successfully")
	fmt.Printf("Database user: %s\n", conf) // Add this line to debug the configuration values
	return conf
}

func setupDatabase(conf *config.Config) *bun.DB {
	return database.NewDB(conf)
}

func Execute() {
	conf := readConfigFileFromEnv()
	fmt.Println(conf)
	appLogger := logger.NewApiLogger(conf)
	appLogger.InitLogger()
	appLogger.Info("Starting the API Server")

	// Initialize and start servers
	redisClient, err := setupRedis(conf, appLogger)
	if err != nil {
		appLogger.Fatal("Shutting down server ", "error: ", err)
	}

	db := setupDatabase(conf)
	sess := setupAWS(conf, appLogger)
	e := setupHTTPServer(conf, db, sess, appLogger, redisClient)
	s := setupGRPCServer(conf, db)

	// Start servers
	startServers(e, conf.HTTP, s, conf.GRPC, db)

	// Graceful shutdown
	gracefulShutdown(e, s, appLogger)
}

func startServers(e *echo.Echo, httpConfig config.HTTP, s *grpc.Server, grpcConfig config.GRPC, db *bun.DB) {
	go startHTTPServer(e, httpConfig)
	go startGRPCServer(s, grpcConfig, db)
}

func gracefulShutdown(e *echo.Echo, s *grpc.Server, appLogger logger.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	shutdownServers(ctx, e, s, appLogger)
}

func shutdownServers(ctx context.Context, e *echo.Echo, s *grpc.Server, appLogger logger.Logger) {
	appLogger.Info("Shutting down gRPC server...")
	s.GracefulStop()
	appLogger.Info("gRPC server stopped!")

	appLogger.Info("Shutting down HTTP server...")
	if err := e.Shutdown(ctx); err != nil {
		appLogger.Fatal(err)
	}
	appLogger.Info("HTTP server stopped!")
}
