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
	IsNegative bool
	Data       []int //todo сделать приватными: пока чисто для тестов
}

func SetBytes(b []byte) *BigDigit {
	result := &BigDigit{}
	str := string(b)
	if str[0] == '-' {
		result.IsNegative = true
		str = str[1:]
	}

	var countBucket = int(math.Ceil(float64(len(str)) / float64(POW)))
	result.Data = make([]int, countBucket)

	bucket := 0
	for i := len(str); i > 0; i -= POW {
		var slice string
		if i < POW {
			slice = str[0:i]
		} else {
			slice = str[i-POW : i]
		}
		result.Data[bucket], _ = strconv.Atoi(slice)
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
	if d1.IsNegative != d2.IsNegative {
		if d1.IsNegative {
			return -1
		}
		return 1
	}
	if len(d1.Data) != len(d2.Data) {
		if len(d1.Data) < len(d2.Data) {
			if d1.IsNegative {
				return 1
			}
			return -1
		}
		if d1.IsNegative {
			return -1
		}
		return 1
	}
	for i := len(d1.Data) - 1; i >= 0; i-- {
		if d1.Data[i] < d2.Data[i] {
			if d1.IsNegative {
				return 1
			}
			return -1
		} else if d1.Data[i] > d2.Data[i] {
			if d1.IsNegative {
				return -1
			}
			return 1
		}
	}
	return 0
}

func CmpModule(d1 *BigDigit, d2 *BigDigit) int8 {
	if len(d1.Data) != len(d2.Data) {
		if len(d1.Data) < len(d2.Data) {
			return -1
		}
		return 1
	}
	for i := len(d1.Data) - 1; i >= 0; i-- {
		if d1.Data[i] < d2.Data[i] {
			return -1
		} else if d1.Data[i] > d2.Data[i] {
			return 1
		}
	}
	return 0
}

func Sum(d1 *BigDigit, d2 *BigDigit) *BigDigit {
	cmp := CmpModule(d1, d2)
	var data []int
	var isNegative bool
	if d1.IsNegative == d2.IsNegative {
		if d1.IsNegative {
			isNegative = true
		}
		return &BigDigit{Data: sumNotNegative(d1.Data, d2.Data), IsNegative: isNegative}
	}
	switch cmp {
	case 0:
		return &BigDigit{Data: []int{0}, IsNegative: false}
	case 1:
		if d1.IsNegative {
			isNegative = true
		}
		data = subNotNegative(d1.Data, d2.Data)
	case -1:
		if d2.IsNegative {
			isNegative = true
		}
		data = subNotNegative(d2.Data, d1.Data)
	}
	return &BigDigit{Data: data, IsNegative: isNegative}
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
	var IsNegative bool
	if cmp == 0 && d1.IsNegative == d2.IsNegative {
		return &BigDigit{Data: []int{0}, IsNegative: false}
	}
	switch {
	case cmp >= 0:
		if d1.IsNegative == d2.IsNegative {
			data = subNotNegative(d1.Data, d2.Data)
		} else {
			data = sumNotNegative(d1.Data, d2.Data)
		}
		if d1.IsNegative {
			IsNegative = true
		}
	case cmp == -1:
		if d1.IsNegative == d2.IsNegative {
			data = subNotNegative(d2.Data, d1.Data)
		} else {
			data = sumNotNegative(d2.Data, d1.Data)
		}
		if d2.IsNegative {
			IsNegative = true
		}
	}
	return &BigDigit{Data: data, IsNegative: IsNegative}
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
	return RemoveBeginZero(res)
}
