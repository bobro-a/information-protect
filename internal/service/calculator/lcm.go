package calculator

import (
	"information-protect/internal/model"
	"information-protect/internal/utils"
)

func (c calculator) LCM(a, b model.BigDigit) model.BigDigit {
	zero := model.BigDigit{Data: []int64{0}, IsNegative: false}

	// Берем модули чисел
	aAbs := c.Abs(a)
	bAbs := c.Abs(b)

	// Если одно из чисел 0 → LCM = 0
	if utils.CompareModule(aAbs.Data, zero.Data) == 0 || utils.CompareModule(bAbs.Data, zero.Data) == 0 {
		return zero
	}

	// LCM(a, b) = |a*b| / GCD(a, b)
	product := c.Mul(aAbs, bAbs) // модуль произведения
	gcd := c.GCD(aAbs, bAbs)     // GCD по модулю

	quotient, _ := c.Div(product, gcd) // деление по модулю
	quotient.IsNegative = false        // LCM всегда ≥ 0

	return quotient
}
