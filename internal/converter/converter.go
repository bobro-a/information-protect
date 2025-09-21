package converter

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"information-protect/internal/model"
)

// FromFile считывает большое число из файла и конвертирует в BigDigit
func FromFile(filePath string) (model.BigDigit, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return model.BigDigit{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var digits []int64
	isNegative := false
	firstChar := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Проверяем знак только для первого непробельного символа
		for _, ch := range line {
			if firstChar {
				if ch == '-' {
					isNegative = true
					firstChar = false
					continue
				} else if ch == '+' {
					firstChar = false
					continue
				}
				firstChar = false
			}

			if ch >= '0' && ch <= '9' {
				n, _ := strconv.ParseInt(string(ch), 10, 64)
				digits = append(digits, n)
			}
		}
	}

	// Если нет цифр, считаем 0
	if len(digits) == 0 {
		digits = []int64{0}
		isNegative = false
	}

	// Переворачиваем, чтобы младший разряд был первым
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return model.BigDigit{Data: digits, IsNegative: isNegative}, nil
}

// ToFile записывает BigDigit в файл
func ToFile(filePath string, num model.BigDigit) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if num.IsNegative {
		_, _ = writer.WriteString("-")
	}

	// Старший разряд первым
	for i := len(num.Data) - 1; i >= 0; i-- {
		_, err := writer.WriteString(strconv.FormatInt(num.Data[i], 10))
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
