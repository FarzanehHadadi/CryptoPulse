-- +goose Up
CREATE TABLE IF NOT EXISTS candles
    (
        id          INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        symbol      VARCHAR(255) NOT NULL                           ,
        open_price  DECIMAL(18, 8) NOT NULL                         ,
        close_price DECIMAL(18, 8) NOT NULL                         ,
        low_price   DECIMAL(18, 8) NOT NULL                         ,
        high_price  DECIMAL(18, 8) NOT NULL                         ,
                    interval VARCHAR(10) NOT NULL                   ,
        volume      DECIMAL(18, 8) NOT NULL                         ,
        created_at  TIMESTAMP NOT NULL DEFAULT NOW()                ,
        open_time   TIMESTAMP NOT NULL                              ,
        close_time  TIMESTAMP NOT NULL                              ,
        is_closed   BOOLEAN NOT NULL DEFAULT TRUE                   ,
        CONSTRAINT uq_candle UNIQUE ( symbol, interval, open_time )
    )
;
-- +goose Down
DROP TABLE IF EXISTS candles;