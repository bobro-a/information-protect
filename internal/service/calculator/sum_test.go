package calculator

import (
	"testing"

	"information-protect/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name string
		a, b model.BigDigit
		want model.BigDigit
	}{
		{
			name: "positive + positive",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: false},
			want: model.BigDigit{Data: []int64{8}, IsNegative: false},
		},
		{
			name: "negative + negative",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: true},
			want: model.BigDigit{Data: []int64{8}, IsNegative: true},
		},
		{
			name: "positive + negative, |a| > |b|",
			a:    model.BigDigit{Data: []int64{7}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: true},
			want: model.BigDigit{Data: []int64{2}, IsNegative: false},
		},
		{
			name: "positive + negative, |b| > |a|",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: true},
			want: model.BigDigit{Data: []int64{2}, IsNegative: true},
		},
		{
			name: "negative + positive, |a| > |b|",
			a:    model.BigDigit{Data: []int64{7}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: false},
			want: model.BigDigit{Data: []int64{2}, IsNegative: true},
		},
		{
			name: "negative + positive, |b| > |a|",
			a:    model.BigDigit{Data: []int64{3}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: false},
			want: model.BigDigit{Data: []int64{2}, IsNegative: false},
		},
		{
			name: "opposite numbers equal",
			a:    model.BigDigit{Data: []int64{5}, IsNegative: true},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: false},
			want: model.BigDigit{Data: []int64{0}, IsNegative: false},
		},
		{
			name: "zero + positive",
			a:    model.BigDigit{Data: []int64{0}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: false},
			want: model.BigDigit{Data: []int64{5}, IsNegative: false},
		},
		{
			name: "zero + negative",
			a:    model.BigDigit{Data: []int64{0}, IsNegative: false},
			b:    model.BigDigit{Data: []int64{5}, IsNegative: true},
			want: model.BigDigit{Data: []int64{5}, IsNegative: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Sum(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "Sum() result should match expected value")
		})
	}
}
