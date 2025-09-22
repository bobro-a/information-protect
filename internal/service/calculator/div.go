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

	for utils.CompareModuleWithRemoveZeros(remainder.Data, bAbs.Data) >= 0 {
		tmp := bAbs
		mult := model.BigDigit{Data: []int64{1}, IsNegative: false}

		for {
			tmpShift := c.Mul(tmp, model.BigDigit{Data: []int64{base}, IsNegative: false})
			if utils.CompareModuleWithRemoveZeros(tmpShift.Data, remainder.Data) > 0 {
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

func (c calculator) divForLcm(a, b model.BigDigit) model.BigDigit {
	if len(b.Data) == 1 && b.Data[0] == 0 {
		panic("деление на ноль")
	}
	if len(a.Data) == 1 && a.Data[0] == 0 {
		return model.BigDigit{Data: []int64{0}, IsNegative: false}
	}

	A := c.Abs(a)
	B := c.Abs(b)

	if utils.CompareModule(A.Data, B.Data) < 0 {
		return model.BigDigit{Data: []int64{0}, IsNegative: false}
	}

	qMSB := make([]int64, 0, len(A.Data))
	remainder := model.BigDigit{Data: []int64{}, IsNegative: false}

	// идём от старших разрядов к младшим
	for i := len(A.Data) - 1; i >= 0; i-- {
		// переносим разряд в остаток
		remainder.Data = append([]int64{A.Data[i]}, remainder.Data...)
		utils.RemoveLeadingZeros(remainder.Data)

		if utils.CompareModule(remainder.Data, B.Data) < 0 {
			qMSB = append(qMSB, 0)
			continue
		}

		// прикидка цифры частного
		lr := utils.LeadingDigits(remainder.Data, 2)
		lb := utils.LeadingDigits(B.Data, 2)
		q := int64(0)
		if lb == 0 {
			q = 9
		} else {
			q = lr / lb
			if q < 1 {
				q = 1
			} else if q > 9 {
				q = 9
			}
		}

		// бинпоиск корректной q
		low, high := q, int64(9)
		prod := c.Mul(B, model.BigDigit{Data: []int64{q}, IsNegative: false})
		for utils.CompareModule(prod.Data, remainder.Data) > 0 {
			high = q - 1
			q = (low + high) / 2
			if q < 0 {
				q = 0
			}
			prod = c.Mul(B, model.BigDigit{Data: []int64{q}, IsNegative: false})
		}
		for {
			if q == 9 {
				break
			}
			test := c.Mul(B, model.BigDigit{Data: []int64{q + 1}, IsNegative: false})
			if utils.CompareModule(test.Data, remainder.Data) <= 0 {
				q++
				prod = test
				continue
			}
			break
		}

		qMSB = append(qMSB, q)
		remainder = c.Sub(remainder, prod)
	}

	// убираем ведущие нули
	k := 0
	for k < len(qMSB)-1 && qMSB[k] == 0 {
		k++
	}
	qMSB = qMSB[k:]

	quotient := model.BigDigit{IsNegative: false}
	for i := len(qMSB) - 1; i >= 0; i-- {
		quotient.Data = append(quotient.Data, qMSB[i])
	}
	utils.RemoveLeadingZeros(quotient.Data)
	if len(quotient.Data) == 0 {
		quotient.Data = []int64{0}
		quotient.IsNegative = false
	}
	return quotient
}
