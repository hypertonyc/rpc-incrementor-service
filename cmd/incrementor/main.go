package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hypertonyc/rpc-incrementor-service/internal/api/grpc"
	"github.com/hypertonyc/rpc-incrementor-service/internal/config"
	"github.com/hypertonyc/rpc-incrementor-service/internal/database"
	"github.com/hypertonyc/rpc-incrementor-service/internal/logger"
	"github.com/hypertonyc/rpc-incrementor-service/internal/repository"
	"github.com/hypertonyc/rpc-incrementor-service/internal/service"
)

func main() {
	// Load configuration from environment variables
	appConfig := config.LoadConfigFromEnv()

	// Initialize logger (folder for log files, log level)
	logger.InitLogger(appConfig)

	var incrementRepo repository.IncrementRepository

	if len(os.Args) == 2 && os.Args[1] == "--inmemory" {
		incrementRepo = repository.NewInMemoryIncrementRepository()
		logger.Logger.Info("server will start with in-memory storage (for debug usage only!)")
	} else {
		// Run db migrations
		err := database.Migrate(appConfig)
		if err != nil {
			log.Fatal("failed to run database migrations:", err)
		}

		// Create the database connection
		db, err := sql.Open("postgres", appConfig.PgConUrl)
		if err != nil {
			log.Fatal("failed to open database connection:", err)
		}
		defer db.Close()

		// Initialize repository with the database connection
		incrementRepo = repository.NewPostgresIncrementRepository(db, context.Background())
	}

	incrementService := service.NewIncrementService(incrementRepo)

	// Create the gRPC server
	grpcServer := grpc.NewServer(appConfig, incrementService)

	// Start the gRPC server
	grpcServer.Start()

	// Listen for OS signals to perform a graceful shutdown
	waitForShutdownSignals()

	grpcServer.Stop()
	logger.Logger.Info("graceful shutdown...")
}

func waitForShutdownSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
