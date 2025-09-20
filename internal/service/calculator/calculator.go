package calculator

import (
	"information-protect/internal/model"
	"information-protect/internal/utils"
)

type Calculator interface {
	Sum(a, b model.BigDigit) model.BigDigit
	Sub(a, b model.BigDigit) model.BigDigit
	Mul(a, b model.BigDigit) model.BigDigit
	Pow(d model.BigDigit, n int) model.BigDigit
	Div(a, b model.BigDigit) (quotient, remainder model.BigDigit)
	GCD(a, b model.BigDigit) model.BigDigit //НОК
	LCM(a, b model.BigDigit) model.BigDigit //НОД
}

type calculator struct {
	base int64 // = 10^pow
	pow  int64
}

func NewCalculator(pow int64) Calculator {
	var base int64 = 1
	for i := int64(0); i < pow; i++ {
		base *= 10
	}
	return &calculator{
		base: base,
		pow:  pow,
	}
}

func (c calculator) GCD(a, b model.BigDigit) model.BigDigit {
	zero := model.BigDigit{Data: []int64{0}, IsNegative: false}
	for utils.CompareModule(b.Data, zero.Data) != 0 {
		_, remainder := c.Div(a, b)
		remainder.Data = utils.RemoveLeadingZeros(remainder.Data)
		a, b = b, remainder
	}
	a.IsNegative = false
	return a
}

func (c calculator) LCM(a, b model.BigDigit) model.BigDigit {
	zero := model.BigDigit{Data: []int64{0}, IsNegative: false}
	if utils.CompareModule(a.Data, zero.Data) == 0 || utils.CompareModule(b.Data, zero.Data) == 0 {
		return zero
	}
	// LCM(a, b) = |a*b| / GCD(a, b)
	product := c.Mul(a, b)
	gcd := c.GCD(a, b)
	quotient, _ := c.Div(product, gcd)
	quotient.IsNegative = false
	return quotient
}
