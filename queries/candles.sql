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