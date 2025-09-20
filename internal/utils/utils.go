package utils

// RemoveLeadingZeros удаляет ведущие нули (старший разряд) в BigDigit
func RemoveLeadingZeros(arr []int64) []int64 {
	// ищем индекс последнего ненулевого элемента с конца (старший разряд)
	last := len(arr) - 1
	for last > 0 && arr[last] == 0 {
		last--
	}
	return arr[:last+1]
}

// CompareModule сравнивает два числа по модулю
// возвращает 1 если d1>d2, -1 если d1<d2, 0 если равны
func CompareModule(d1, d2 []int64) int8 {
	d1 = RemoveLeadingZeros(d1)
	d2 = RemoveLeadingZeros(d2)

	if len(d1) > len(d2) {
		return 1
	} else if len(d1) < len(d2) {
		return -1
	}

	// сравниваем старший разряд к младшему
	for i := len(d1) - 1; i >= 0; i-- {
		if d1[i] > d2[i] {
			return 1
		} else if d1[i] < d2[i] {
			return -1
		}
	}
	return 0
}

// Автоматический расчет base по двум слайсам
func AutoBase(d1, d2 []int64) int64 {
	maxElem := int64(0)
	for _, v := range d1 {
		if v > maxElem {
			maxElem = v
		}
	}
	for _, v := range d2 {
		if v > maxElem {
			maxElem = v
		}
	}

	base := int64(1)
	for base <= maxElem {
		base *= 10
	}
	return base
}
