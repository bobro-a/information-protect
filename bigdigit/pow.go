package bigdigit

func isOdd(data int) bool {
	return data%2 == 1
}

// Алгоритм Карацубы
func Mult(d1 *BigDigit, d2 *BigDigit) *BigDigit {
	len1, len2 := len(d1.data), len(d2.data)

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
	a := &BigDigit{data: d1.data[N:]}
	b := &BigDigit{data: d1.data[:N]}
	c := &BigDigit{data: d1.data[N:]}
	d := &BigDigit{data: d1.data[:N]}

	ac := Mult(a, c)
	bd := Mult(b, d)
	ab_cd := Sub(Sub(Mult(Sum(a, b), Sum(c, d)), ac), bd)

	ac.data = append(make([]int, 2*N), ac.data...)
	ab_cd.data = append(make([]int, N), ab_cd.data...)

	return &BigDigit{data: Sum(Sum(ac, bd), ab_cd).data}
}

// умножение столбиком
func simpleMult(d1 *BigDigit, d2 *BigDigit) *BigDigit {
	resData := make([]int, len(d1.data)+len(d2.data))

	for i, digit1 := range d1.data {
		carry := 0
		for j, digit2 := range d2.data {
			res := digit1*digit2 + carry + resData[i+j]
			resData[i+j] = res % 10
			carry = res / 10
		}
		if carry > 0 {
			resData[i+len(d2.data)] += carry
		}
	}

	return &BigDigit{data: resData}
}

func Pow(d1 *BigDigit, degree int) *BigDigit {
	resData := make([]int, len(d1.data))
	resData = append(resData, 1)
	base := &BigDigit{data: append([]int{}, d1.data...)}

	for degree > 0 {
		if isOdd(degree) {
			resData = Mult(&BigDigit{data: resData}, base).data
		}
		base = Mult(base, base)
		degree /= 2
	}

	return &BigDigit{data: resData}
}

//// pow возводит BigDigit в степень degree.
//func (d *BigDigit) pow(degree int) *BigDigit {
//	// Создаем "единицу" для инициализации результата
//	res := NewBigDigit("1")
//
//	// Создаем копию основания
//	base := &BigDigit{data: append([]int{}, d.data...)}
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
