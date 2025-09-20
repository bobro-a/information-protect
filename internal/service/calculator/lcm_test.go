package calculator

import (
	"testing"

	"information-protect/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestLCM(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name string
		a, b model.BigDigit
		want model.BigDigit
	}{
		{
			name: "simple lcm",
			a:    model.BigDigit{Data: []int64{4}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{6}, IsNegative: false},
			want: model.BigDigit{Data: []int64{2, 1}, IsNegative: false},
		},
		{
			name: "lcm with negative number",
			a:    model.BigDigit{Data: []int64{4}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{6}, IsNegative: false},
			want: model.BigDigit{Data: []int64{2, 1}, IsNegative: false},
		},
		{
			name: "lcm with zero",
			a:    model.BigDigit{Data: []int64{0}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{6}, IsNegative: false},
			want: model.BigDigit{Data: []int64{0}, IsNegative: false},
		},
		{
			name: "both zero",
			a:    model.BigDigit{Data: []int64{0}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{0}, IsNegative: false},
			want: model.BigDigit{Data: []int64{0}, IsNegative: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.LCM(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "LCM() result should match expected value")
		})
	}
}
