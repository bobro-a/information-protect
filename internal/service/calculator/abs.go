package calculator

import "information-protect/internal/model"

func (c calculator) Abs(a model.BigDigit) model.BigDigit {
	return model.BigDigit{Data: a.Data,
		IsNegative: false}
}
