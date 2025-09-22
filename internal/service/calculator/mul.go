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

	len1, len2 := len(aAbs.Data), len(bAbs.Data)

	if len1 < 15 && len2 < 15 {
		res = c.simpleMul(aAbs, bAbs, base)
		res.Data = utils.RemoveLeadingZeros(res.Data)
		res.IsNegative = a.IsNegative != b.IsNegative
		return res
	}

	n := len1
	if len2 > n {
		n = len2
	}
	if n%2 != 0 {
		n++
	}

	N := n / 2
	a1 := model.BigDigit{Data: aAbs.Data[N:]}
	b1 := model.BigDigit{Data: aAbs.Data[:N]}
	c1 := model.BigDigit{Data: bAbs.Data[N:]}
	d := model.BigDigit{Data: bAbs.Data[:N]}

	ac := c.Mul(a1, c1)
	bd := c.Mul(b1, d)
	ab_cd := c.Sub(c.Sub(c.Mul(c.Sum(a1, b1), c.Sum(c1, d)), ac), bd)

	ac.Data = append(make([]int64, 2*N), ac.Data...)
	ab_cd.Data = append(make([]int64, N), ab_cd.Data...)

	res.Data = c.Sum(c.Sum(ac, bd), ab_cd).Data

	res.Data = utils.RemoveLeadingZeros(res.Data)

	res.IsNegative = a.IsNegative != b.IsNegative

	return res
}

func (c calculator) simpleMul(a, b model.BigDigit, base int64) (res model.BigDigit) {

	res.Data = make([]int64, len(a.Data)+len(b.Data))

	for i := 0; i < len(a.Data); i++ {
		var carry int64
		for j := 0; j < len(b.Data); j++ {
			tmp := a.Data[i]*b.Data[j] + res.Data[i+j] + carry
			res.Data[i+j] = tmp % base
			carry = tmp / base
		}
		if carry != 0 {
			res.Data[i+len(b.Data)] += carry
		}
	}

	return res
}
