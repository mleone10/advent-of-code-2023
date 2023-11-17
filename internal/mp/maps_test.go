package mp_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/mp"
)

func TestKeys(t *testing.T) {
	tcs := []struct {
		subj     map[int]int
		expected []int
	}{
		{
			map[int]int{1: 5, 2: 10, 3: 15},
			[]int{1, 2, 3},
		},
	}

	for _, tc := range tcs {
		actual := mp.Keys(tc.subj)
		assert.ArrayEquals(t, tc.expected, actual)
	}
}
