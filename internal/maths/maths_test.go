package maths_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/maths"
)

func TestFactors(t *testing.T) {
	tcs := []struct {
		n        int
		expected []int
	}{
		{
			1,
			[]int{1},
		},
		{
			4,
			[]int{1, 2, 4},
		},
		{
			150,
			[]int{1, 2, 3, 5, 6, 10, 15, 25, 30, 50, 75, 150},
		},
		{
			100000,
			[]int{1, 2, 4, 5, 8, 10, 16, 20, 25, 32, 40, 50, 80, 100, 125, 160, 200, 250, 400, 500, 625, 800, 1000, 1250, 2000, 2500, 3125, 4000, 5000, 6250, 10000, 12500, 20000, 25000, 50000, 100000},
		},
	}

	for _, tc := range tcs {
		actual := maths.Factors(tc.n)
		assert.ArrayEquals(t, tc.expected, actual)
	}
}
