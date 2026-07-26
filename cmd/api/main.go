package main

import (
	"context"
	"cryptoPulse/internal/backfill"
	"cryptoPulse/internal/collector"
	"cryptoPulse/internal/config"
	"cryptoPulse/internal/database"
	"cryptoPulse/internal/db"
	"cryptoPulse/internal/exchanges/binanceEx"
	"cryptoPulse/internal/logger"
	"cryptoPulse/internal/repository"
	"cryptoPulse/internal/scheduler"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	//***********load env*************//
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found, using system environment variables")
	}
	//***********init logger*************//
	logConfig := config.GetLoggerConfig()
	logger.Init(logConfig)
	logger.Info(
		"application started",
		"env", cfg.App.Environment,
	)
	//***********init database*************//
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
	//***********init repository*************//
	queries := db.New(pool)
	repo := repository.NewRepository(queries, pool)

	binanceClient := binanceEx.NewClient("", "")
	collectorService := collector.NewService(repo, binanceClient)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	backfillService := backfill.NewService(binanceClient, repo.Candles, binanceClient.Capabilities().MaxCandleLimit)

	jobs := scheduler.JobsFromConfig(cfg.Scheduler) // or inline map
	for _, job := range jobs {
		if err := backfillService.Run(ctx, job.CandleRequest); err != nil {
			logger.Error("Backfill failed", "job", job.Name, "error", err)
		}
	}
	sched := scheduler.NewScheduler(collectorService, jobs, cfg.Scheduler.Workers)
	if err := sched.Start(ctx); err != nil && err != context.Canceled {
		logger.Error("scheduler stopped", "error", err)
	}

}
