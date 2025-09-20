package calculator

import (
	"testing"

	"information-protect/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name        string
		a           model.BigDigit
		b           model.BigDigit
		wantQ       model.BigDigit
		wantR       model.BigDigit
		expectPanic bool
	}{
		{
			name:  "simple division",
			a:     model.BigDigit{Data: []int64{10}, IsNegative: false},
			b:     model.BigDigit{Data: []int64{2}, IsNegative: false},
			wantQ: model.BigDigit{Data: []int64{5}, IsNegative: false},
			wantR: model.BigDigit{Data: []int64{0}, IsNegative: false},
		},
		{
			name:  "division with remainder",
			a:     model.BigDigit{Data: []int64{10}, IsNegative: false},
			b:     model.BigDigit{Data: []int64{3}, IsNegative: false},
			wantQ: model.BigDigit{Data: []int64{3}, IsNegative: false},
			wantR: model.BigDigit{Data: []int64{1}, IsNegative: false},
		},
		{
			name:  "division negative dividend",
			a:     model.BigDigit{Data: []int64{10}, IsNegative: true},
			b:     model.BigDigit{Data: []int64{2}, IsNegative: false},
			wantQ: model.BigDigit{Data: []int64{5}, IsNegative: true},
			wantR: model.BigDigit{Data: []int64{0}, IsNegative: true},
		},
		{
			name:  "division negative divisor",
			a:     model.BigDigit{Data: []int64{10}, IsNegative: false},
			b:     model.BigDigit{Data: []int64{2}, IsNegative: true},
			wantQ: model.BigDigit{Data: []int64{5}, IsNegative: true},
			wantR: model.BigDigit{Data: []int64{0}, IsNegative: false},
		},
		{
			name:  "division both negative",
			a:     model.BigDigit{Data: []int64{10}, IsNegative: true},
			b:     model.BigDigit{Data: []int64{2}, IsNegative: true},
			wantQ: model.BigDigit{Data: []int64{5}, IsNegative: false},
			wantR: model.BigDigit{Data: []int64{0}, IsNegative: true},
		},
		{
			name:        "division by zero",
			a:           model.BigDigit{Data: []int64{10}, IsNegative: false},
			b:           model.BigDigit{Data: []int64{0}, IsNegative: false},
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				assert.Panics(t, func() {
					calc.Div(tt.a, tt.b)
				}, "expected panic on division by zero")
			} else {
				q, r := calc.Div(tt.a, tt.b)
				assert.Equal(t, tt.wantQ, q, "quotient should match")
				assert.Equal(t, tt.wantR, r, "remainder should match")
			}
		})
	}
}
