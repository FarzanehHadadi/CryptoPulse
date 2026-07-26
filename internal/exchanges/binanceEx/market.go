package binanceEx

import (
	"context"
	"cryptoPulse/internal/domain"
	"cryptoPulse/internal/exchange"
)

func (c *Client) GetCandles(ctx context.Context, request exchange.CandleRequest) ([]domain.Candle, error) {
	nkc := c.client.NewKlinesService().
		Symbol(request.Symbol).
		Interval(request.Interval).
		Limit(request.Limit)
	if request.StartTime != nil {
		nkc.StartTime(request.StartTime.UnixMilli())
	}

	if request.EndTime != nil {
		nkc.EndTime(request.EndTime.UnixMilli())
	}

	klines, err := nkc.
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
func (b *Client) Capabilities() exchange.Capabilities {
	return exchange.Capabilities{
		MaxCandleLimit: 1000,
	}
}
