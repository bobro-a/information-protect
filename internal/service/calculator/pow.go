package calculator

import "information-protect/internal/model"

func (c calculator) Pow(d model.BigDigit, n int) (res model.BigDigit) {
	if n == 0 {
		return model.BigDigit{Data: []int64{1}, IsNegative: false}
	}

	if n < 0 {
		panic("отрицательные степени пока не поддерживаются")
	}

	res = model.BigDigit{Data: []int64{1}, IsNegative: false}
	base := d

	for n > 0 {
		if n%2 == 1 {
			res = c.Mul(res, base)
		}
		base = c.Mul(base, base)
		n /= 2
	}

	return
}
