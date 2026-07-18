-- name: GetSymbols :many
SELECT *
FROM symbols;

-- name: GetSymbolByID :one
SELECT *
FROM symbols
WHERE id = $1;