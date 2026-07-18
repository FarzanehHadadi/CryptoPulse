package config

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Logger   LoggerConfig
}

type AppConfig struct {
	Name        string
	Environment string
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type RedisConfig struct {
	Host string
	Port int
}

type LoggerConfig struct {
	Environment string
	ServiceName string
	Version     string
	LogLevel    string
}
