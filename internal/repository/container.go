package repository

import (
	"cryptoPulse/internal/db"
)

func NewRepository(queries *db.Queries) *Repository {
	return &Repository{
		Symbols: NewSymbolsRepository(queries),
	}
}
