package calculator

import (
	"testing"

	"information-protect/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name string
		a, b model.BigDigit
		want model.BigDigit
	}{
		{
			name: "positive minus smaller positive",
			a:    model.BigDigit{Data: []int64{5}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{3}, IsNegative: false},
			want: model.BigDigit{Data: []int64{2}, IsNegative: false},
		},
		{
			name: "positive minus larger positive",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: false},
			want: model.BigDigit{Data: []int64{2}, IsNegative: true},
		},
		{
			name: "negative minus positive",
			a:    model.BigDigit{Data: []int64{4}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{3}, IsNegative: false},
			want: model.BigDigit{Data: []int64{7}, IsNegative: true},
		},
		{
			name: "positive minus negative",
			a:    model.BigDigit{Data: []int64{4}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{3}, IsNegative: true},
			want: model.BigDigit{Data: []int64{7}, IsNegative: false},
		},
		{
			name: "negative minus negative, |a| > |b|",
			a:    model.BigDigit{Data: []int64{5}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{3}, IsNegative: true},
			want: model.BigDigit{Data: []int64{2}, IsNegative: true},
		},
		{
			name: "negative minus negative, |b| > |a|",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: true},
			want: model.BigDigit{Data: []int64{2}, IsNegative: false},
		},
		{
			name: "equal positive numbers",
			a:    model.BigDigit{Data: []int64{7}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{7}, IsNegative: false},
			want: model.BigDigit{Data: []int64{0}, IsNegative: false},
		},
		{
			name: "equal negative numbers",
			a:    model.BigDigit{Data: []int64{7}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{7}, IsNegative: true},
			want: model.BigDigit{Data: []int64{0}, IsNegative: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Sub(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "Sub() result should match expected value")
		})
	}
}
