CREATE TABLE indicators (
    id BIGSERIAL PRIMARY KEY,

    symbol TEXT NOT NULL,
    interval TEXT NOT NULL,

    indicator_name TEXT NOT NULL,

    period INTEGER,

    candle_time TIMESTAMPTZ NOT NULL,

    value NUMERIC(20,8) NOT NULL,

    created_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT uq_indicator UNIQUE (symbol, interval, indicator_name, period, candle_time)
);