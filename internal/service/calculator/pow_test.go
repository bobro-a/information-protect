package calculator

import (
	"testing"

	"information-protect/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name        string
		base        model.BigDigit
		exp         int
		want        model.BigDigit
		expectPanic bool
	}{
		{
			name: "positive exponent",
			base: model.BigDigit{Data: []int64{2}, IsNegative: false},
			exp:  3,
			want: model.BigDigit{Data: []int64{8}, IsNegative: false},
		},
		{
			name: "zero exponent",
			base: model.BigDigit{Data: []int64{5}, IsNegative: false},
			exp:  0,
			want: model.BigDigit{Data: []int64{1}, IsNegative: false},
		},
		{
			name: "negative base odd exponent",
			base: model.BigDigit{Data: []int64{2}, IsNegative: true},
			exp:  3,
			want: model.BigDigit{Data: []int64{8}, IsNegative: true},
		},
		{
			name: "negative base even exponent",
			base: model.BigDigit{Data: []int64{2}, IsNegative: true},
			exp:  4,
			want: model.BigDigit{Data: []int64{6, 1}, IsNegative: false},
		},
		{
			name:        "negative exponent should panic",
			base:        model.BigDigit{Data: []int64{2}, IsNegative: false},
			exp:         -1,
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				assert.Panics(t, func() {
					calc.Pow(tt.base, tt.exp)
				}, "Pow() with negative exponent should panic")
			} else {
				got := calc.Pow(tt.base, tt.exp)
				assert.Equal(t, tt.want, got, "Pow() result should match expected value")
			}
		})
	}
}
