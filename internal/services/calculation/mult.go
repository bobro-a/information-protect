package calculation

import "infp01/internal/model"

func (c calculator) Mult() model.BigDigit {
	len1, len2 := len(c.d1.Data), len(c.d2.Data)

	if len1 < 15 && len2 < 15 {
		return simpleMult(c.d1, c.d2)
	}

	n := len1
	if len2 > n {
		n = len2
	}
	if n%2 != 0 {
		n++
	}

	N := n / 2
	a := model.BigDigit{Data: c.d1.Data[N:]}
	b := model.BigDigit{Data: c.d1.Data[:N]}
	c2 := model.BigDigit{Data: c.d1.Data[N:]}
	d := model.BigDigit{Data: c.d1.Data[:N]}

	p1 := calculator{d1: a, d2: c2}
	p2 := calculator{d1: b, d2: d}

	p1 = p1.Mult()
	p2 = p2.Mult()
	ab_cd := Sub(Sub(Mult(Sum(a, b), Sum(c2, d)), ac), bd)

	ac.Data = append(make([]int, 2*N), ac.Data...)
	ab_cd.Data = append(make([]int, N), ab_cd.Data...)

	return &model.BigDigit{Data: Sum(Sum(ac, bd), ab_cd).Data}
}

// умножение столбиком
func simpleMult(c.d1 *BigDigit, c.d2 *BigDigit) *BigDigit {
	resData := make([]int, len(c.d1.Data)+len(c.d2.Data))

	for i, digit1 := range c.d1.Data {
		carry := 0
		for j, digit2 := range c.d2.Data {
			res := digit1*digit2 + carry + resData[i+j]
			resData[i+j] = res % 10
			carry = res / 10
		}
		if carry > 0 {
			resData[i+len(c.d2.Data)] += carry
		}
	}

	return &BigDigit{Data: resData}
}

func Pow(c.d1 *BigDigit, degree int) *BigDigit {
	resData := make([]int, len(c.d1.Data))
	resData = append(resData, 1)
	base := &BigDigit{Data: append([]int{}, c.d1.Data...)}

	for degree > 0 {
		if isOdd(degree) {
			resData = Mult(&BigDigit{Data: resData}, base).Data
		}
		base = Mult(base, base)
		degree /= 2
	}

	return &BigDigit{Data: resData}
}

//// pow возводит BigDigit в степень degree.
//func (d *BigDigit) pow(degree int) *BigDigit {
//	// Создаем "единицу" для инициализации результата
//	res := NewBigDigit("1")
//
//	// Создаем копию основания
//	base := &BigDigit{Data: append([]int{}, d.Data...)}
//
//	for degree > 0 {
//		// Если степень нечетная, умножаем результат на основание
//		if degree%2 == 1 {
//			res = res.Mult(base)
//		}
//		// Возводим основание в квадрат
//		base = base.Mult(base)
//		// Делим степень на 2
//		degree /= 2
//	}
//
//	return res
//}
