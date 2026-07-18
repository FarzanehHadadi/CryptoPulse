package database

import (
	"context"
	"cryptoPulse/internal/utils"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type dbConfig struct {
	Host            string
	Port            string
	Username        string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifeTime time.Duration
	ConnMaxIdleTime time.Duration
}

func DefaultConfig() dbConfig {
	return dbConfig{
		Host:            utils.GetEnv("DB_HOST", "localhost"),
		Port:            utils.GetEnv("DB_PORT", "5434"),
		Username:        utils.GetEnv("DB_USER", "postgres"),
		Password:        utils.GetEnv("DB_PASSWORD", "postgres"),
		DBName:          utils.GetEnv("DB_NAME", "cryptopulse"),
		SSLMode:         utils.GetEnv("DB_SSLMODE", "disable"),
		MaxOpenConn:     25,
		MaxIdleConn:     10,
		ConnMaxLifeTime: time.Minute * 5,
		ConnMaxIdleTime: time.Minute * 10,
	}
}

func NewPostgresConfiguration(cfg *dbConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	config.MaxConns = int32(cfg.MaxOpenConn)
	config.MinConns = int32(cfg.MaxIdleConn)

	config.MaxConnLifetime = cfg.ConnMaxLifeTime
	config.MaxConnIdleTime = cfg.ConnMaxIdleTime

	pool, err := pgxpool.NewWithConfig(
		context.Background(),
		config,
	)

	if err != nil {
		return nil, err
	}

	return pool, nil

}
