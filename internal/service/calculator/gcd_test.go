package calculator

import (
	"testing"

	"information-protect/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name string
		a, b model.BigDigit
		want model.BigDigit
	}{
		{
			name: "simple gcd",
			a:    model.BigDigit{Data: []int64{48}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{18}, IsNegative: false},
			want: model.BigDigit{Data: []int64{6}, IsNegative: false},
		},
		{
			name: "gcd negative numbers",
			a:    model.BigDigit{Data: []int64{48}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{18}, IsNegative: true},
			want: model.BigDigit{Data: []int64{6}, IsNegative: false},
		},
		{
			name: "gcd one zero",
			a:    model.BigDigit{Data: []int64{0}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{18}, IsNegative: false},
			want: model.BigDigit{Data: []int64{18}, IsNegative: false},
		},
		{
			name: "gcd both zero",
			a:    model.BigDigit{Data: []int64{0}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{0}, IsNegative: false},
			want: model.BigDigit{Data: []int64{0}, IsNegative: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.GCD(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "GCD() result should match expected value")
		})
	}
}
