package godecimal

// Add 加 (d+d2)
func (d Decimal) Add(d2 Decimal) Decimal {
	mul := New()
	mul.floatValue.Add(d.floatValue, d2.floatValue)
	return mul
}

// Sub 减 (d-d2)
func (d Decimal) Sub(d2 Decimal) Decimal {
	mul := New()
	mul.floatValue.Sub(d.floatValue, d2.floatValue)
	return mul
}

// Mul 乘 (d*d2)
func (d Decimal) Mul(d2 Decimal) Decimal {
	mul := New()
	mul.floatValue.Mul(d.floatValue, d2.floatValue)
	return mul
}

// Quo 除 (d/d2)
func (d Decimal) Quo(d2 Decimal) Decimal {
	mul := New()
	mul.floatValue.Quo(d.floatValue, d2.floatValue)
	return mul
}
