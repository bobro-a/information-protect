package calculator

import (
	"fmt"

	"information-protect/internal/model"
	"information-protect/internal/utils"
)

func (c calculator) Mul(a, b model.BigDigit) (res model.BigDigit) {
	lenA, lenB := len(a.Data), len(b.Data)
	temp := make([]int64, lenA+lenB)

	for i := 0; i < lenA; i++ {
		var remain int64
		for j := 0; j < lenB; j++ {
			tmp := a.Data[i]*b.Data[j] + temp[i+j] + remain
			temp[i+j] = tmp % c.base
			remain = tmp / c.base
		}
		if remain > 0 {
			temp[i+lenB] += remain
		}
	}

	res = model.BigDigit{
		Data:       utils.RemoveLeadingZeros(temp),
		IsNegative: a.IsNegative != b.IsNegative,
	}
	fmt.Println("test mul")
	return
}
