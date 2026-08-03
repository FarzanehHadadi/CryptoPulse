package mapper

import (
	"cryptoPulse/internal/db"
	"cryptoPulse/internal/domain"
	"cryptoPulse/internal/utils"

	"github.com/jackc/pgx/v5/pgtype"
)

func DbCandleToDomain(candle db.Candle) (domain.Candle, error) {
	open, err := utils.NumericToDecimal(candle.OpenPrice)
	if err != nil {
		return domain.Candle{}, err
	}
	high, err := utils.NumericToDecimal(candle.HighPrice)
	if err != nil {
		return domain.Candle{}, err
	}
	low, err := utils.NumericToDecimal(candle.LowPrice)
	if err != nil {
		return domain.Candle{}, err
	}
	closePrice, err := utils.NumericToDecimal(candle.ClosePrice)
	if err != nil {
		return domain.Candle{}, err
	}
	volume, err := utils.NumericToDecimal(candle.Volume)
	if err != nil {
		return domain.Candle{}, err
	}

	return domain.Candle{
		ID:        candle.ID,
		Symbol:    candle.Symbol,
		Interval:  candle.Interval,
		Open:      open,
		High:      high,
		Low:       low,
		Close:     closePrice,
		Volume:    volume,
		OpenTime:  candle.OpenTime.Time,
		CloseTime: candle.CloseTime.Time,
	}, nil
}

func DbCandlesToDomain(candles []db.Candle) ([]domain.Candle, error) {
	result := make([]domain.Candle, 0, len(candles))
	for _, candle := range candles {
		mapped, err := DbCandleToDomain(candle)
		if err != nil {
			return nil, err
		}
		result = append(result, mapped)
	}
	return result, nil
}

func DomainCandleToCreateParams(c domain.Candle) db.CreateCandlesParams {
	return db.CreateCandlesParams{
		Symbol:     c.Symbol,
		OpenPrice:  utils.DecimalToNumeric(c.Open),
		ClosePrice: utils.DecimalToNumeric(c.Close),
		LowPrice:   utils.DecimalToNumeric(c.Low),
		HighPrice:  utils.DecimalToNumeric(c.High),
		Interval:   c.Interval,
		Volume:     utils.DecimalToNumeric(c.Volume),
		OpenTime:   pgtype.Timestamp{Time: c.OpenTime, Valid: true},
		CloseTime:  pgtype.Timestamp{Time: c.CloseTime, Valid: true},
		IsClosed:   true,
	}
}

func DomainCandlesToCreateParams(candles []domain.Candle) []db.CreateCandlesParams {
	params := make([]db.CreateCandlesParams, 0, len(candles))
	for _, c := range candles {
		params = append(params, DomainCandleToCreateParams(c))
	}
	return params
}
