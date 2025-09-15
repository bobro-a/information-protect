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
	data       []int64
}

func SetBytes(b []byte) *BigDigit {
	result := &BigDigit{}
	str := string(b)
	if str[0] == '-' {
		result.isNegative = true
		str = str[1:]
	}

	var countBucket = int(math.Ceil(float64(len(str)) / float64(POW)))
	result.data = make([]int64, countBucket)

	bucket := 0
	for i := len(str); i > 0; i -= POW {
		var slice string
		if i < POW {
			slice = str[0:i]
		} else {
			slice = str[i-POW : i]
		}
		result.data[bucket], _ = strconv.ParseInt(slice, 10, 64)
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
func Sum(d1 *BigDigit, d2 *BigDigit) {
	
}
