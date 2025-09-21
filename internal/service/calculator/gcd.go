package calculator

import (
	"information-protect/internal/model"
	"information-protect/internal/utils"
)

func (c calculator) GCD(a, b model.BigDigit) model.BigDigit {
	zero := model.BigDigit{Data: []int64{0}, IsNegative: false}
	a = c.Abs(a)
	b = c.Abs(b)

	for utils.CompareModule(b.Data, zero.Data) != 0 {
		_, remainder := c.Div(a, b)
		remainder.Data = utils.RemoveLeadingZeros(remainder.Data)
		remainder.IsNegative = false // на всякий случай
		a, b = b, remainder
	}
	a.IsNegative = false
	return a
}
