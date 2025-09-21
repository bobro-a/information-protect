package calculator

import (
	"information-protect/internal/model"
)

type Calculator interface {
	Sum(a, b model.BigDigit) model.BigDigit
	Sub(a, b model.BigDigit) model.BigDigit
	Mul(a, b model.BigDigit) model.BigDigit
	Pow(d model.BigDigit, n int) model.BigDigit
	Div(a, b model.BigDigit) (quotient, remainder model.BigDigit)
	GCD(a, b model.BigDigit) model.BigDigit //НОД
	LCM(a, b model.BigDigit) model.BigDigit //НОК
	Abs(a model.BigDigit) model.BigDigit
}

type calculator struct{}

func NewCalculator() Calculator {
	return &calculator{}
}
