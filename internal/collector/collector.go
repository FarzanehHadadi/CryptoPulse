package collector

import (
	"cryptoPulse/internal/exchange"
	"cryptoPulse/internal/repository"
)

type Service struct {
	repo     repository.Repository
	exchange exchange.Exchange
}

func NewService(repo *repository.Repository, exchange exchange.Exchange) *Service {
	return &Service{
		repo:     *repo,
		exchange: exchange,
	}
}
