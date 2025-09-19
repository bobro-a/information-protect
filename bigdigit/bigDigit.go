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

func sumNotNegative(d1 *BigDigit, d2 *BigDigit) *BigDigit {
	i, j, remains := 0, 0, 0
	var BASE = int(math.Pow10(POW))
	size := max(len(d1.data), len(d2.data))
	var res BigDigit
	res.data = make([]int, size)
	for ; i < len(d1.data) && j < len(d2.data); i, j = i+1, j+1 {
		res.data[i] = (d1.data[i] + d2.data[j] + remains) % BASE
		remains = (d1.data[i] + d2.data[j] + remains) / BASE
	}
	for ; i < len(d1.data); i++ {
		res.data[i] = (d1.data[i] + remains) % BASE
		remains = (d1.data[i] + remains) / BASE
	}
	for ; j < len(d2.data); j++ {
		res.data[j] = (d2.data[j] + remains) % BASE
		remains = (d2.data[j] + remains) / BASE
	}
	if remains != 0 {
		res.data = append(res.data, remains)
	}
	return &res
}

func subNotNegative(d1 *BigDigit, d2 *BigDigit) *BigDigit {
	var res BigDigit
	var BASE = int(math.Pow10(POW))
	i, j, loan := 0, 0, 0
	size := min(len(d1.data), len(d2.data))
	res.data = make([]int, size)

	for ; i < len(d1.data) && j < len(d2.data); i, j = i+1, j+1 {
		sub := d1.data[i] - loan - d2.data[j]
		if sub >= 0 {
			res.data[i] = sub
			loan = 0
		} else {
			res.data[i] = sub + BASE
			loan = 0
		}
	}
	//todo
	return &res
}
