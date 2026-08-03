package indicator

import "github.com/shopspring/decimal"

type Result struct {
	Period int
	Values []decimal.Decimal // one SMA per window; aligns with candles[N-1:]
}
