package repository

import (
	"context"
	"cryptoPulse/internal/db"
	"cryptoPulse/internal/domain"
	"cryptoPulse/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type indicatorRepo struct {
	db    *pgxpool.Pool
	query *db.Queries
}

func NewIndicatorsRepository(queries *db.Queries, db *pgxpool.Pool) *indicatorRepo {
	return &indicatorRepo{
		db:    db,
		query: queries,
	}
}

func (r *indicatorRepo) CreateIndicators(ctx context.Context, indicators []domain.IndicatorResult) error {
	if len(indicators) == 0 {
		return nil
	}

	params := make([]db.CreateIndicatorsParams, 0, len(indicators))
	for _, indicator := range indicators {
		params = append(params, db.CreateIndicatorsParams{
			Symbol:        indicator.Symbol,
			Interval:      indicator.Interval,
			IndicatorName: indicator.Name,
			Period:        pgtype.Int4{Int32: int32(indicator.Period), Valid: indicator.Period > 0},
			CandleTime:    pgtype.Timestamptz{Time: indicator.CandleTime, Valid: true},
			Value:         utils.DecimalToNumeric(indicator.Value),
		})
	}

	batch := r.query.CreateIndicators(ctx, params)
	var batchErr error
	batch.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	return batchErr
}
