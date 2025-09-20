package calculation

import (
	"infp01/internal/model"
	"infp01/internal/utils"
)

func (c calculator) Sum() model.BigDigit {
	cmp := utils.CompareModule(c.d1.Data, c.d2.Data)
	var data []int
	var isNegative bool
	if c.d1.IsNegative == c.d2.IsNegative {
		if c.d1.IsNegative {
			isNegative = true
		}
		return model.BigDigit{Data: sumNotNegative(c.d1.Data, c.d2.Data), IsNegative: isNegative}
	}
	switch cmp {
	case 0:
		return model.BigDigit{Data: []int{0}, IsNegative: false}
	case 1:
		if c.d1.IsNegative {
			isNegative = true
		}
		data = subNotNegative(c.d1.Data, c.d2.Data)
	case -1:
		if c.d2.IsNegative {
			isNegative = true
		}
		data = subNotNegative(c.d2.Data, c.d1.Data)
	}
	return model.BigDigit{Data: data, IsNegative: isNegative}
}
