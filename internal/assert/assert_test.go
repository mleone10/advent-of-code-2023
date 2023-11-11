package assert_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
)

func AssertArrayEquals(t *testing.T) {
	tcs := []struct {
		a, b     []int
		expected bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
			false,
		},
	}

	for _, tc := range tcs {
		actual := assert.ArrayEquals(tc.a, tc.b)
		if actual != tc.expected {
			t.Errorf("expected [%v], got [%v]", tc.expected, actual)
		}
	}
}
