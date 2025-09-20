package calculator

import (
	"information-protect/internal/model"
	"information-protect/internal/utils"
)

func (c calculator) Mul(a, b model.BigDigit) (res model.BigDigit) {
	// Работаем с модулями чисел
	aAbs := c.Abs(a) // возвращает BigDigit с IsNegative=false
	bAbs := c.Abs(b)

	// Автоматически определяем базу
	base := utils.AutoBase(aAbs.Data, bAbs.Data)

	// Результат может быть длины len(a)+len(b)
	res.Data = make([]int64, len(aAbs.Data)+len(bAbs.Data))

	for i := 0; i < len(aAbs.Data); i++ {
		var carry int64
		for j := 0; j < len(bAbs.Data); j++ {
			tmp := aAbs.Data[i]*bAbs.Data[j] + res.Data[i+j] + carry
			res.Data[i+j] = tmp % base
			carry = tmp / base
		}
		if carry != 0 {
			res.Data[i+len(bAbs.Data)] += carry
		}
	}

	res.Data = utils.RemoveLeadingZeros(res.Data)

	res.IsNegative = a.IsNegative != b.IsNegative

	return res
}
