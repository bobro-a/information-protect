package calculator

import (
	"information-protect/internal/model"
	"information-protect/internal/utils"
)

func (c calculator) LCM(a, b model.BigDigit) model.BigDigit {
	zero := model.BigDigit{Data: []int64{0}, IsNegative: false}

	// Если одно из чисел 0 → LCM = 0
	if utils.CompareModule(a.Data, zero.Data) == 0 || utils.CompareModule(b.Data, zero.Data) == 0 {
		return zero
	}

	// |a| и |b|
	aAbs := c.Abs(a)
	bAbs := c.Abs(b)

	// gcd = GCD(|a|, |b|)
	gcd := c.GCD(aAbs, bAbs)

	t := c.divForLcm(aAbs, gcd)

	res := c.Mul(t, bAbs)

	res.IsNegative = false

	return res
}
