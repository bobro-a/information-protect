package generator

import (
	"math/rand"
	"time"

	"information-protect/internal/model"
)

// GenerateBigNumber генерирует большое число длиной size цифр
func GenerateBigNumber(size int) model.BigDigit {
	rand.Seed(time.Now().UnixNano())
	data := make([]int64, size)
	for i := 0; i < size; i++ {
		data[i] = int64(rand.Intn(10)) // цифры 0..9
	}
	return model.BigDigit{Data: data, IsNegative: false}
}
