-- name: CreateIndicators :batchexec
INSERT INTO indicators (
    symbol,
    interval,
    indicator_name,
    period,
    candle_time,
    value
)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (symbol, interval, indicator_name, period, candle_time)
DO NOTHING;
