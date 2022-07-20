package godecimal

import (
	"math/big"
)

// String 输出 string
func (d Decimal) String() string {
	return d.floatValue.String()
}

// Float64 输出 float64
func (d Decimal) Float64() float64 {
	rat, _ := new(big.Rat).SetString(d.String())
	f, _ := rat.Float64()
	return f
}

// MoneyFloat64 货币 float64
func (d Decimal) MoneyFloat64() float64 {
	rat, _ := new(big.Rat).SetString(d.floatValue.Text('f', 2))
	f, _ := rat.Float64()
	return f
}
