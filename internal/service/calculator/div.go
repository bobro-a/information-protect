package calculator

import (
	"information-protect/internal/model"
	"information-protect/internal/utils"
)

func (c calculator) Div(a, b model.BigDigit) (quotient, remainder model.BigDigit) {
	base := utils.AutoBase(a.Data, b.Data)
	if len(b.Data) == 1 && b.Data[0] == 0 {
		panic("деление на ноль")
	}

	quotient = model.BigDigit{Data: []int64{0}, IsNegative: false}
	remainder = c.Abs(a) // используем модуль
	bAbs := c.Abs(b)

	for utils.CompareModule(remainder.Data, bAbs.Data) >= 0 {
		tmp := bAbs
		mult := model.BigDigit{Data: []int64{1}, IsNegative: false}

		for {
			tmpShift := c.Mul(tmp, model.BigDigit{Data: []int64{base}, IsNegative: false})
			if utils.CompareModule(tmpShift.Data, remainder.Data) > 0 {
				break
			}
			tmp = tmpShift
			mult = c.Mul(mult, model.BigDigit{Data: []int64{base}, IsNegative: false})
		}

		remainder = c.Sub(remainder, tmp)
		quotient = c.Sum(quotient, mult)
	}

	quotient.IsNegative = a.IsNegative != b.IsNegative
	remainder.IsNegative = a.IsNegative

	return quotient, remainder
}
