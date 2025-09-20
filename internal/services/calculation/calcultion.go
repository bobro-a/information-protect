package calculation

import (
	"infp01/internal/model"
	"infp01/internal/utils"
	"math"
)

const (
	POW = 9
)

type Calculator interface {
	Sum() model.BigDigit
}

type calculator struct {
	d1 model.BigDigit
	d2 model.BigDigit
}

func New(d1 model.BigDigit, d2 model.BigDigit) Calculator {
	return &calculator{d1: d1, d2: d2}
}

func sumNotNegative(d1 []int, d2 []int) []int {
	i, j, remains := 0, 0, 0
	var BASE = int(math.Pow10(POW))
	size := max(len(d1), len(d2))
	res := make([]int, size)
	for ; i < len(d1) && j < len(d2); i, j = i+1, j+1 {
		res[i] = (d1[i] + d2[j] + remains) % BASE
		remains = (d1[i] + d2[j] + remains) / BASE
	}
	for ; i < len(d1); i++ {
		res[i] = (d1[i] + remains) % BASE
		remains = (d1[i] + remains) / BASE
	}
	for ; j < len(d2); j++ {
		res[j] = (d2[j] + remains) % BASE
		remains = (d2[j] + remains) / BASE
	}
	if remains != 0 {
		res = append(res, remains)
	}
	return res
}

func subNotNegative(largerNum []int, smallerNum []int) []int {
	var BASE = int(math.Pow10(POW))
	i, loan := 0, 0
	res := make([]int, len(largerNum))

	for ; i < len(largerNum); i++ {
		sub := largerNum[i] - loan
		if i < len(smallerNum) {
			sub -= smallerNum[i]
		}

		if sub >= 0 {
			loan = 0
		} else {
			sub += BASE
			loan = 1
		}
		res[i] = sub
	}
	return utils.RemoveBeginZero(res)
}
