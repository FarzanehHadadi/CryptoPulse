package config

type Config struct {
	App       AppConfig
	Server    ServerConfig
	Database  DatabaseConfig
	Redis     RedisConfig
	Logger    LoggerConfig
	Scheduler SchedulerConfig
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
type SchedulerConfig struct {
	Workers int
	Jobs    []JobConfig
}
type JobConfig struct {
	Name          string
	CandleRequest CandleRequest
}
type CandleRequest struct {
	Symbol   string
	Interval string
	Limit    int
}
