package binanceEx

import (
	"cryptoPulse/internal/logger"

	"github.com/adshao/go-binance/v2"
)

type Client struct {
	client *binance.Client
}

func NewClient(apiKey string, secretKey string) *Client {
	logger.Info("Initializing Binance client", "apiKey", apiKey, "secretKey", secretKey)
	client := binance.NewClient("", "")
	// client := binance.NewClient(apiKey, secretKey)
	return &Client{
		client: client,
	}
}
