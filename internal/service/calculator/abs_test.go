package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"information-protect/internal/model"
)

func TestAbs(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name  string
		input model.BigDigit
		want  model.BigDigit
	}{
		{
			name:  "positive number",
			input: model.BigDigit{Data: []int64{123, 456}, IsNegative: false},
			want:  model.BigDigit{Data: []int64{123, 456}, IsNegative: false},
		},
		{
			name:  "negative number",
			input: model.BigDigit{Data: []int64{123, 456}, IsNegative: true},
			want:  model.BigDigit{Data: []int64{123, 456}, IsNegative: false},
		},
		{
			name:  "zero",
			input: model.BigDigit{Data: []int64{0}, IsNegative: true},
			want:  model.BigDigit{Data: []int64{0}, IsNegative: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Abs(tt.input)
			assert.Equal(t, tt.want.IsNegative, got.IsNegative, "IsNegative should match")
			assert.Equal(t, tt.want.Data, got.Data, "Data slices should match")
		})
	}
}
