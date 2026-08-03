package repository

import (
	"context"
	"cryptoPulse/internal/db"
	"cryptoPulse/internal/domain"
	"cryptoPulse/internal/repository/mapper"

	"github.com/jackc/pgx/v5/pgxpool"
)

type candlesRepository struct {
	db    *pgxpool.Pool
	query *db.Queries
}

func NewCandlesRepository(queries *db.Queries, db *pgxpool.Pool) *candlesRepository {
	return &candlesRepository{
		db:    db,
		query: queries,
	}
}

func (r *candlesRepository) CreateCandles(ctx context.Context, candles []domain.Candle) error {
	batch := r.query.CreateCandles(ctx, mapper.DomainCandlesToCreateParams(candles))
	var batchErr error
	batch.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	return batchErr
}

func (r *candlesRepository) GetLastCandleBySymbol(ctx context.Context, symbol string, interval string) (*domain.Candle, error) {
	candle, err := r.query.GetLastCandleBySymbol(ctx, db.GetLastCandleBySymbolParams{
		Symbol:   symbol,
		Interval: interval,
	})
	if err != nil {
		return nil, err
	}
	mapped, err := mapper.DbCandleToDomain(candle)
	if err != nil {
		return nil, err
	}
	return &mapped, nil
}

func (r *candlesRepository) GetCandlesBySymbol(ctx context.Context, symbol string, interval string, count int) ([]domain.Candle, error) {
	if count <= 0 {
		return []domain.Candle{}, nil
	}

	rows, err := r.query.GetCandlesBySymbol(ctx, db.GetCandlesBySymbolParams{
		Symbol:   symbol,
		Interval: interval,
		Limit:    int32(count),
	})
	if err != nil {
		return nil, err
	}

	return mapper.DbCandlesToDomain(rows)
}
