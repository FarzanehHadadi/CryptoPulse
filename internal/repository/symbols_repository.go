package repository

import (
	"context"
	"cryptoPulse/internal/db"
)

type symbolsRepository struct {
	query *db.Queries
}

func NewSymbolsRepository(query *db.Queries) *symbolsRepository {
	return &symbolsRepository{
		query: query,
	}
}

func (s *symbolsRepository) GetSymbols() ([]db.Symbol, error) {
	return s.query.GetSymbols(context.Background())
}
