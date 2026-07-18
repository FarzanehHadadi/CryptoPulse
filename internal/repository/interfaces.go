package repository

import "cryptoPulse/internal/db"

type Repository struct {
	Symbols Symbols
}
type Symbols interface {
	GetSymbols() ([]db.Symbol, error)
}
