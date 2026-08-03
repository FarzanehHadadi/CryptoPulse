-- name: CreateCandles :batchexec
INSERT INTO candles
    (
        symbol     ,
        open_price ,
        close_price,
        low_price  ,
        high_price ,
        interval   ,
        volume     ,
        open_time  ,
        close_time ,
        is_closed
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10
    )
    ON CONFLICT
    (
        symbol  ,
        interval,
        open_time
    )
    DO NOTHING;

-- name: GetLastCandleBySymbol :one
SELECT *
FROM candles
WHERE symbol = $1
  AND interval = $2
ORDER BY open_time DESC
LIMIT 1;

-- name: GetCandlesBySymbol :many
SELECT id, symbol, open_price, close_price, low_price, high_price, interval, volume, created_at, open_time, close_time, is_closed
FROM (
    SELECT *
    FROM candles
    WHERE symbol = $1
      AND interval = $2
    ORDER BY open_time DESC
    LIMIT $3
) AS recent
ORDER BY open_time ASC;