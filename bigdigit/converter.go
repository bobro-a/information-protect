package bigdigit

import (
	"io"
	"math"
	"os"
	"strconv"
)

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

func GetFile(path string, digit *BigDigit) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if digit.IsNegative {
		file.Write([]byte("-"))
	}
	for i := len(digit.Data) - 1; i >= 0; i-- {
		file.Write([]byte(strconv.Itoa(digit.Data[i])))
	}
	return nil
}
