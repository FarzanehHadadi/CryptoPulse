package domain

import (
	"fmt"
	"time"
)

func ParseInterval(interval string) (time.Duration, error) {
	switch interval {
	case "1s":
		return time.Second, nil
	case "1m":
		return time.Minute, nil
	case "3m":
		return 3 * time.Minute, nil
	case "5m":
		return 5 * time.Minute, nil
	case "15m":
		return 15 * time.Minute, nil
	case "30m":
		return 30 * time.Minute, nil
	case "1h":
		return time.Hour, nil
	case "2h":
		return 2 * time.Hour, nil
	case "4h":
		return 4 * time.Hour, nil
	case "1d":
		return 24 * time.Hour, nil
	case "1w":
		return 7 * 24 * time.Hour, nil
	// "1M" ≈ 30 days only if you accept approximation
	default:
		return 0, fmt.Errorf("unsupported interval: %s", interval)
	}
}
