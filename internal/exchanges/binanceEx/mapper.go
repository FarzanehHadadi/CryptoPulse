package binanceEx

import (
	"cryptoPulse/internal/domain"
	"cryptoPulse/internal/logger"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/shopspring/decimal"
)

func mapKlineToCandle(kline *binance.Kline, symbol string, interval string) domain.Candle {
	openPrice, err := decimal.NewFromString(kline.Open)
	if err != nil {
		logger.Error("invalid open price", "error", err)
		return domain.Candle{}
	}
	highPrice, err := decimal.NewFromString(kline.High)
	if err != nil {
		logger.Error("invalid high price", "error", err)
		return domain.Candle{}
	}
	lowPrice, err := decimal.NewFromString(kline.Low)
	if err != nil {
		logger.Error("invalid low price", "error", err)
		return domain.Candle{}
	}
	closePrice, err := decimal.NewFromString(kline.Close)
	if err != nil {
		logger.Error("invalid close price", "error", err)
		return domain.Candle{}
	}
	volume, err := decimal.NewFromString(kline.Volume)
	if err != nil {
		logger.Error("invalid volume", "error", err)
		return domain.Candle{}
	}
	return domain.Candle{
		Symbol:    symbol,
		Interval:  interval,
		Open:      openPrice,
		High:      highPrice,
		Low:       lowPrice,
		Close:     closePrice,
		Volume:    volume,
		OpenTime:  time.Unix(kline.OpenTime/1000, 0),
		CloseTime: time.Unix(kline.CloseTime/1000, 0),
	}
}
