package repository

import (
	"context"
	"cryptoPulse/internal/db"
	"cryptoPulse/internal/domain"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
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
	params := make([]db.CreateCandlesParams, 0, len(candles))
	for _, c := range candles {
		params = append(params, db.CreateCandlesParams{
			Symbol:     c.Symbol,
			OpenPrice:  decimalToNumeric(c.Open),
			ClosePrice: decimalToNumeric(c.Close),
			LowPrice:   decimalToNumeric(c.Low),
			HighPrice:  decimalToNumeric(c.High),
			Interval:   c.Interval,
			Volume:     decimalToNumeric(c.Volume),
			OpenTime:   pgtype.Timestamp{Time: c.OpenTime, Valid: true},
			CloseTime:  pgtype.Timestamp{Time: c.CloseTime, Valid: true},
			IsClosed:   true,
		})
	}
	batch := r.query.CreateCandles(ctx, params)
	var batchErr error
	batch.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	return batchErr
}
func decimalToNumeric(d decimal.Decimal) pgtype.Numeric {
	var n pgtype.Numeric
_:
	n.Scan(d.String())
	return n
}
