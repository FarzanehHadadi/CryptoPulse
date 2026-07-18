package main

import (
	"context"
	"cryptoPulse/internal/config"
	"cryptoPulse/internal/database"
	"cryptoPulse/internal/db"
	"cryptoPulse/internal/logger"
	"cryptoPulse/internal/repository"
	"log/slog"

	"github.com/joho/godotenv"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	//load env
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found, using system environment variables")
	}
	//init logger
	logConfig := config.GetLoggerConfig()

	logger.Init(logConfig)
	slog.Info(
		"application started",
		"env", cfg.App.Environment,
	)
	dbConfig := database.DefaultConfig()
	pool, err := database.NewPostgresConfiguration(&dbConfig)
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	logger.Info("Database connected successfully")

	if err := pool.Ping(context.Background()); err != nil {
		panic(err)
	}
	logger.Info("Database pinged successfully")
	queries := db.New(pool)
	repo := repository.NewRepository(queries)
	symbols, err := repo.Symbols.GetSymbols()
	if err != nil {
		panic(err)
	}
	logger.Info("Symbols fetched successfully", "symbols", symbols)

}
