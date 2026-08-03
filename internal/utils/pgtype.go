package utils

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

func DecimalToNumeric(d decimal.Decimal) pgtype.Numeric {
	var n pgtype.Numeric
	_ = n.Scan(d.String())
	return n
}

func NumericToDecimal(n pgtype.Numeric) (decimal.Decimal, error) {
	if !n.Valid {
		return decimal.Zero, nil
	}
	v, err := n.Value()
	if err != nil {
		return decimal.Zero, err
	}
	s, ok := v.(string)
	if !ok {
		return decimal.Zero, nil
	}
	return decimal.NewFromString(s)
}
