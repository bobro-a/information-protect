package bigdigit

import (
	"io"
	"math"
	"os"
	"strconv"
)

const (
	POW = 9
)

type BigDigit struct {
	isNegative bool
	data       []int
}

func SetBytes(b []byte) *BigDigit {
	result := &BigDigit{}
	str := string(b)
	if str[0] == '-' {
		result.isNegative = true
		str = str[1:]
	}

	var countBucket = int(math.Ceil(float64(len(str)) / float64(POW)))
	result.data = make([]int, countBucket)

	bucket := 0
	for i := len(str); i > 0; i -= POW {
		var slice string
		if i < POW {
			slice = str[0:i]
		} else {
			slice = str[i-POW : i]
		}
		result.data[bucket], _ = strconv.Atoi(slice)
		bucket++
	}
	return result
}

func SetFile(path string) (*BigDigit, error) {
	res := &BigDigit{}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return res, err
	}
	filesize := fileInfo.Size()
	file, err := os.Open(path)
	if err != nil {
		return res, err
	}
	defer file.Close()

	buffer := make([]byte, filesize)
	for {
		_, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
	}
	res = SetBytes(buffer)
	return res, nil
}

func CmpDigit(d1 *BigDigit, d2 *BigDigit) int8 {
	if d1.isNegative != d2.isNegative {
		if d1.isNegative {
			return -1
		}
		return 1
	}
	if len(d1.data) != len(d2.data) {
		if len(d1.data) < len(d2.data) {
			if d1.isNegative {
				return 1
			}
			return -1
		}
		if d1.isNegative {
			return -1
		}
		return 1
	}
	for i := len(d1.data) - 1; i >= 0; i-- {
		if d1.data[i] < d2.data[i] {
			if d1.isNegative {
				return 1
			}
			return -1
		} else if d1.data[i] > d2.data[i] {
			if d1.isNegative {
				return -1
			}
			return 1
		}
	}
	return 0
}

func CmpModule(d1 *BigDigit, d2 *BigDigit) int8 {
	if len(d1.data) != len(d2.data) {
		if len(d1.data) < len(d2.data) {
			return -1
		}
		return 1
	}
	for i := len(d1.data) - 1; i >= 0; i-- {
		if d1.data[i] < d2.data[i] {
			return -1
		} else if d1.data[i] > d2.data[i] {
			return 1
		}
	}
	return 0
}

func Sum(d1 *BigDigit, d2 *BigDigit) *BigDigit {
	return &BigDigit{data: make([]int, 0)}
	//todo
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

func Sub(d1 *BigDigit, d2 *BigDigit) *BigDigit {
	cmp := CmpModule(d1, d2)
	var data []int
	var isNegative bool
	if cmp == 0 && d1.isNegative == d2.isNegative {
		return &BigDigit{data: []int{0}, isNegative: false}
	}
	switch {
	case cmp >= 0:
		if d1.isNegative == d2.isNegative {
			data = subNotNegative(d1.data, d2.data)
		} else {
			data = sumNotNegative(d1.data, d2.data)
		}
		if d1.isNegative {
			isNegative = true
		}
	case cmp == -1:
		if d1.isNegative == d2.isNegative {
			data = subNotNegative(d2.data, d1.data)
		} else {
			data = sumNotNegative(d2.data, d1.data)
		}
		if d2.isNegative {
			isNegative = true
		}
	}
	return &BigDigit{data: data, isNegative: isNegative}
}

func subNotNegative(largerNum []int, smallerNum []int) []int {
	var BASE = int(math.Pow10(POW))
	i, j, loan := 0, 0, 0
	res := make([]int, len(largerNum))

	for ; i < len(largerNum); i++ {
		sub := largerNum[i] - loan
		if i < len(smallerNum) {
			sub -= smallerNum[j]
		}

		if sub >= 0 {
			loan = 0
		} else {
			sub += BASE
			loan = 1
		}
		res[i] = sub
	}
	lastIndex := len(res) - 1
	for ; lastIndex > 0 && res[lastIndex] == 0; lastIndex-- {
	}
	return res[:lastIndex+1]
}
