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
	}

	for _, tc := range tcs {
		actual := maths.Factors(tc.n)
		assert.ArrayEquals(t, tc.expected, actual)
	}
}
