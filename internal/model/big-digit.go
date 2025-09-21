package model

type BigDigit struct {
	IsNegative bool    `json:"is_negative"`
	Data       []int64 `json:"data"`
}
