package calculation

func (c calculator) Mult() model.BigDigit {
	len1, len2 := len(d1.Data), len(d2.Data)

	if len1 < 15 && len2 < 15 {
		return simpleMult(d1, d2)
	}

	n := len1
	if len2 > n {
		n = len2
	}
	if n%2 != 0 {
		n++
	}

	N := n / 2
	a := &BigDigit{Data: d1.Data[N:]}
	b := &BigDigit{Data: d1.Data[:N]}
	c := &BigDigit{Data: d1.Data[N:]}
	d := &BigDigit{Data: d1.Data[:N]}

	ac := Mult(a, c)
	bd := Mult(b, d)
	ab_cd := Sub(Sub(Mult(Sum(a, b), Sum(c, d)), ac), bd)

	ac.Data = append(make([]int, 2*N), ac.Data...)
	ab_cd.Data = append(make([]int, N), ab_cd.Data...)

	return &BigDigit{Data: Sum(Sum(ac, bd), ab_cd).Data}
}

// умножение столбиком
func simpleMult(d1 *BigDigit, d2 *BigDigit) *BigDigit {
	resData := make([]int, len(d1.Data)+len(d2.Data))

	for i, digit1 := range d1.Data {
		carry := 0
		for j, digit2 := range d2.Data {
			res := digit1*digit2 + carry + resData[i+j]
			resData[i+j] = res % 10
			carry = res / 10
		}
		if carry > 0 {
			resData[i+len(d2.Data)] += carry
		}
	}

	return &BigDigit{Data: resData}
}

func Pow(d1 *BigDigit, degree int) *BigDigit {
	resData := make([]int, len(d1.Data))
	resData = append(resData, 1)
	base := &BigDigit{Data: append([]int{}, d1.Data...)}

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
