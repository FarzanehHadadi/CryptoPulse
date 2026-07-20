package binanceEx

import (
	"context"
	"cryptoPulse/internal/domain"
	"cryptoPulse/internal/exchange"
)

func (c *Client) GetCandles(ctx context.Context, request exchange.CandleRequest) ([]domain.Candle, error) {

	klines, err := c.client.NewKlinesService().
		Symbol(request.Symbol).
		Interval(request.Interval).
		Limit(request.Limit).
		Do(ctx)

	if err != nil {

		return nil, err
	}

	candles := make([]domain.Candle, 0, len(klines))
	for _, kline := range klines {
		candles = append(candles, mapKlineToCandle(kline, request.Symbol, request.Interval))
	}
	return candles, nil
}
