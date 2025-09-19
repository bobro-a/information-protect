package tests

import (
	"laba/bigdigit"
	"slices"
	"testing"
)

func TestSub(t *testing.T) {
	type data struct {
		pathD1 string
		pathD2 string
		result bigdigit.BigDigit
	}
	input := []data{
		{pathD1: "files/digit1.txt",
			pathD2: "files/digit2.txt",
			result: bigdigit.BigDigit{Data: []int{0}, IsNegative: false}},
		{pathD1: "files/digit1.txt",
			pathD2: "files/digit3.txt",
			result: bigdigit.BigDigit{Data: []int{0}, IsNegative: true}},
		{pathD1: "files/digit3.txt",
			pathD2: "files/digit4.txt",
			result: bigdigit.BigDigit{Data: []int{0}, IsNegative: false}},

		{pathD1: "files/digit3.txt",
			pathD2: "files/digit4.txt",
			result: bigdigit.BigDigit{Data: []int{186419754, 975308642}, IsNegative: true}},
	}
	for _, val := range input {
		d1, _ := bigdigit.SetFile(val.pathD1)
		d2, _ := bigdigit.SetFile(val.pathD2)
		res := bigdigit.Sub(d1, d2)
		if !slices.Equal(res.Data, val.result.Data) || res.IsNegative != val.result.IsNegative {
			t.Error("Doesn't match the answer!")
		}
	}
}
