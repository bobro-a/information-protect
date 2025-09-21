package utils

import (
	"reflect"
	"testing"
)

func TestRemoveLeadingZeros(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []int64
	}{
		{"no zeros", []int64{1, 2, 3}, []int64{1, 2, 3}},
		{"trailing zeros", []int64{1, 2, 0, 0}, []int64{1, 2}},
		{"all zeros", []int64{0, 0, 0}, []int64{0}},
		{"single zero", []int64{0}, []int64{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveLeadingZeros(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveLeadingZeros(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCompareModule(t *testing.T) {
	tests := []struct {
		name   string
		d1, d2 []int64
		want   int8
	}{
		{"equal", []int64{1, 2, 3}, []int64{1, 2, 3}, 0},
		{"d1 greater", []int64{1, 2, 4}, []int64{1, 2, 3}, 1},
		{"d2 greater", []int64{1, 2, 3}, []int64{1, 2, 4}, -1},
		{"different length d1 longer", []int64{1, 2, 3, 0}, []int64{1, 2, 3}, 1},
		{"different length d2 longer", []int64{1, 2, 3}, []int64{1, 2, 3, 0}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompareModule(tt.d1, tt.d2)
			if got != tt.want {
				t.Errorf("CompareModule(%v, %v) = %v, want %v", tt.d1, tt.d2, got, tt.want)
			}
		})
	}
}

func TestAutoBase(t *testing.T) {
	tests := []struct {
		name   string
		d1, d2 []int64
		want   int64
	}{
		{"small numbers", []int64{1, 2, 3}, []int64{4, 5, 6}, 10},
		{"larger numbers", []int64{10, 20}, []int64{30, 5}, 100},
		{"single element", []int64{9}, []int64{5}, 10},
		{"all zeros", []int64{0, 0}, []int64{0, 0}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AutoBase(tt.d1, tt.d2)
			if got != tt.want {
				t.Errorf("AutoBase(%v, %v) = %v, want %v", tt.d1, tt.d2, got, tt.want)
			}
		})
	}
}
