package gost

import (
	"math"
	"testing"
)

func eq(a, b float64) bool {
	return math.Abs(b-a) < 0.000001
}

func TestPercentIntegers(t *testing.T) {
	tests := []struct {
		num   int
		denom int
		want  float64
	}{
		{1, 1, 100.0},
		{1, 2, 50.0},
		{2, 3, 66.666666666},
		{97, 100, 97.00000},
	}

	for _, test := range tests {
		got := Percent(test.num, test.denom)

		if !eq(got, test.want) {
			t.Errorf("Percent(%d, %d): Got %f, want %f", test.num, test.denom, got, test.want)
		}
	}
}
