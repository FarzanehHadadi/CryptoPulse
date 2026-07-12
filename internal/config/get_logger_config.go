package config

import "cryptoPulse/internal/utils"

type LoggerConfig struct {
	Environment string
	ServiceName string
	Version     string
	LogLevel    string
}

func GetLoggerConfig() LoggerConfig {
	environment := utils.GetEnv("APP_ENV", "development1")
	serviceName := utils.GetEnv("SERVICE_NAME", "crypto-collector1")
	version := utils.GetEnv("VERSION", "1.0.01")
	logLevel := utils.GetEnv("LOG_LEVEL", "debug1")
	return LoggerConfig{
		Environment: environment,
		ServiceName: serviceName,
		Version:     version,
		LogLevel:    logLevel,
	}
}
