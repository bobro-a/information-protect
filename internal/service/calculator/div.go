package calculator

import (
	"fmt"

	"information-protect/internal/model"
	"information-protect/internal/utils"
)

func (c calculator) Div(a, b model.BigDigit) (quotient, remainder model.BigDigit) {
	if len(b.Data) == 1 && b.Data[0] == 0 {
		panic("деление на ноль")
	}

	quotient = model.BigDigit{Data: []int64{0}, IsNegative: false}
	remainder = a

	for utils.CompareModule(remainder.Data, b.Data) >= 0 {
		tmp := b
		mult := model.BigDigit{Data: []int64{1}, IsNegative: false}

		// Увеличиваем tmp в степени base пока tmp*base <= remainder
		for {
			tmpShift := c.Mul(tmp, model.BigDigit{Data: []int64{c.base}, IsNegative: false})
			if utils.CompareModule(tmpShift.Data, remainder.Data) > 0 {
				break
			}
			tmp = tmpShift
			mult = c.Mul(mult, model.BigDigit{Data: []int64{c.base}, IsNegative: false})
		}

		remainder = c.Sub(remainder, tmp)
		quotient = c.Sum(quotient, mult)
	}

	quotient.IsNegative = a.IsNegative != b.IsNegative
	remainder.IsNegative = a.IsNegative
	fmt.Println("test div")
	return
}
