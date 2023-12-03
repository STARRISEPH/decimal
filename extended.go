package decimal

import "math/big"

func NewFrom100(value int64) Decimal {
	return Decimal{
		value: big.NewInt(value),
		exp:   0,
	}.Div(NewFromInt(100))
}

func (d Decimal) Int100() int64 {
	scaledD := d.rescale(-2)
	return scaledD.value.Int64()
}

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
