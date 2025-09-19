package bigdigit

var (
	zero = BigDigit{Data: []int{0}}
	one  = BigDigit{Data: []int{1}}
)

func (b BigDigit) Div(d *BigDigit) (*BigDigit, *BigDigit) { //result and remainder
	cmp := CmpModule(&b, d)

	if cmp == -1 {
		return &zero, &b
	}
	var isNegative = false
	if d.IsNegative != b.IsNegative {
		isNegative = true
	}
	if cmp == 0 {
		return NewBigDigit(one.Data, isNegative), &zero
	}
	var res, rem BigDigit

	for {
		sub := subNotNegative(b.Data, d.Data)
		if CmpModule(&BigDigit{Data: sub}, &zero) >= 0 {
			res.Inc()
			rem = *NewBigDigit(sub, false)
		} else {
			break
		}
	}
	return &res, &rem
}

func (b *BigDigit) Inc() {
	b.Data = sumNotNegative(b.Data, one.Data)
}
