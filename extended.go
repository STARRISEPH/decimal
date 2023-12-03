package decimal

import "math/big"

func NewFrom10k(value int64) Decimal {
	return Decimal{
		value: big.NewInt(value),
		exp:   0,
	}.Div(NewFromInt(10000))
}

func (d Decimal) Int10k() int64 {
	scaledD := d.rescale(-4)
	return scaledD.value.Int64()
}
