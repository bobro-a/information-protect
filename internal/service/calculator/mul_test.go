package calculator

import (
	"testing"

	"information-protect/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestMul(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name string
		a, b model.BigDigit
		want model.BigDigit
	}{
		{
			name: "positive numbers",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{4}, IsNegative: false},
			want: model.BigDigit{Data: []int64{2, 1}, IsNegative: false},
		},
		{
			name: "negative times positive",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{4}, IsNegative: false},
			want: model.BigDigit{Data: []int64{2, 1}, IsNegative: true},
		},
		{
			name: "positive times negative",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{4}, IsNegative: true},
			want: model.BigDigit{Data: []int64{2, 1}, IsNegative: true},
		},
		{
			name: "negative times negative",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{4}, IsNegative: true},
			want: model.BigDigit{Data: []int64{2, 1}, IsNegative: false},
		},
		{
			name: "multiplication by zero",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{0}, IsNegative: false},
			want: model.BigDigit{Data: []int64{0}, IsNegative: false},
		},
		{
			name: "zero times negative",
			a:    model.BigDigit{Data: []int64{0}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: true},
			want: model.BigDigit{Data: []int64{0}, IsNegative: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Mul(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "Mul() result should match expected value")
		})
	}
}
