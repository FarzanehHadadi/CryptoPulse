-- +goose Up

CREATE TABLE symbols (
    id BIGSERIAL PRIMARY KEY,
    exchange VARCHAR(50) NOT NULL,
    symbol VARCHAR(50) NOT NULL,
    base_asset VARCHAR(20),
    quote_asset VARCHAR(20),
    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down

DROP TABLE symbols;