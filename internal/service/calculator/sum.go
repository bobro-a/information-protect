package calculator

import (
	"information-protect/internal/model"
	"information-protect/internal/utils"
)

func (c calculator) Sum(d1, d2 model.BigDigit) (res model.BigDigit) {
	// Если числа одного знака, просто складываем
	if d1.IsNegative == d2.IsNegative {
		res.Data = c.sumNotNegative(d1.Data, d2.Data)
		res.IsNegative = d1.IsNegative
		return
	}

	// Если числа разных знаков, делаем вычитание: больший по модулю минус меньший
	cmp := utils.CompareModule(d1.Data, d2.Data)
	switch cmp {
	case 0:
		// Результат 0
		res.Data = []int64{0}
		res.IsNegative = false
	case 1:
		// |d1| > |d2|
		res.Data = c.subNotNegative(d1.Data, d2.Data)
		res.IsNegative = d1.IsNegative
	case -1:
		// |d2| > |d1|
		res.Data = c.subNotNegative(d2.Data, d1.Data)
		res.IsNegative = d2.IsNegative
	}
	return res
}

func (c calculator) sumNotNegative(d1 []int64, d2 []int64) []int64 {
	base := utils.AutoBase(d1, d2)
	i, j := 0, 0
	var remains int64

	size := max(len(d1), len(d2))
	res := make([]int64, size)

	for ; i < len(d1) && j < len(d2); i, j = i+1, j+1 {
		sum := d1[i] + d2[j] + remains
		res[i] = sum % base
		remains = sum / base
	}

	for ; i < len(d1); i++ {
		sum := d1[i] + remains
		res[i] = sum % base
		remains = sum / base
	}

	for ; j < len(d2); j++ {
		sum := d2[j] + remains
		res[j] = sum % base
		remains = sum / base
	}

	if remains != 0 {
		res = append(res, remains)
	}

	return res
}
