package main

import (
	"context"
	"cryptoPulse/internal/config"
	"cryptoPulse/internal/logger"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	//init logger
	logConfig := config.GetLoggerConfig()
	logger.Init(logConfig)
	logger.Info("Hello, World!")
	logger.Debug("Debug message")
	logger.Warn("Warn message")
	logger.Error("Error message")
	test := logger.WithContext(context.WithValue(context.Background(), "trace_id", "1234567890"))
	test.Info("Test message")
	test.Debug("Test debug message")
	test.Warn("Test warn message")
	test.Error("Test error message")
}
