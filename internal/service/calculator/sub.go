package calculator

import (
	"information-protect/internal/model"
	"information-protect/internal/utils"
)

func (c calculator) Sub(a, b model.BigDigit) (res model.BigDigit) {
	cmp := utils.CompareModule(a.Data, b.Data)

	// Если числа равны по модулю и имеют одинаковый знак → результат 0
	if cmp == 0 && a.IsNegative == b.IsNegative {
		res.Data = []int64{0}
		res.IsNegative = false
		return
	}

	switch {
	case cmp > 0: // |a| > |b|
		if a.IsNegative == b.IsNegative {
			res.Data = c.subNotNegative(a.Data, b.Data)
			res.IsNegative = a.IsNegative
		} else {
			res.Data = c.sumNotNegative(a.Data, b.Data)
			res.IsNegative = a.IsNegative
		}
	case cmp < 0: // |b| > |a|
		if a.IsNegative == b.IsNegative {
			res.Data = c.subNotNegative(b.Data, a.Data)
			res.IsNegative = !a.IsNegative // знак меняется
		} else {
			res.Data = c.sumNotNegative(a.Data, b.Data)
			res.IsNegative = a.IsNegative
		}
	}

	return
}

func (c calculator) subNotNegative(largerNum []int64, smallerNum []int64) (res []int64) {
	var loan int64
	res = make([]int64, len(largerNum))
	base := c.base

	for i := 0; i < len(largerNum); i++ {
		sub := largerNum[i] - loan
		if i < len(smallerNum) {
			sub -= smallerNum[i]
		}

		if sub >= 0 {
			loan = 0
		} else {
			sub += base
			loan = 1
		}
		res[i] = sub
	}

	return utils.RemoveLeadingZeros(res)
}
