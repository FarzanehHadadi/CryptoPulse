package repository

import (
	"cryptoPulse/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRepository(queries *db.Queries, db *pgxpool.Pool) *Repository {
	return &Repository{
		Symbols: NewSymbolsRepository(queries),
		Candles: NewCandlesRepository(queries, db),
	}
}
