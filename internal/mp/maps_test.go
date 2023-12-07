package mp_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/mp"
)

func TestKeysAndValues(t *testing.T) {
	tcs := []struct {
		subj           map[int]int
		expectedKeys   []int
		expectedValues []int
	}{
		{
			map[int]int{1: 5, 2: 10, 3: 15},
			[]int{1, 2, 3},
			[]int{5, 10, 15},
		},
	}

	for _, tc := range tcs {
		actualKeys := mp.Keys(tc.subj)
		actualValues := mp.Values(tc.subj)
		assert.ArrayEquals(t, tc.expectedKeys, actualKeys)
		assert.ArrayEquals(t, tc.expectedValues, actualValues)
	}
}

func TestMerge(t *testing.T) {
	mapA := map[int]int{1: 2, 3: 4, 5: 6}
	mapB := map[int]int{7: 8, 9: 10}
	mapC := map[int]int{11: 12, 13: 14, 15: 16, 1: 2}

	actual := mp.Merge(mapA, mapB, mapC)
	assert.Equals(t, 8, len(actual))
}

func TestRuneCount(t *testing.T) {
	tcs := []struct {
		input string
		rcs   map[rune]int
	}{
		{
			"foobar",
			map[rune]int{'f': 1, 'o': 2, 'b': 1, 'a': 1, 'r': 1},
		},
		{
			"mississippi",
			map[rune]int{'m': 1, 'i': 4, 's': 4, 'p': 2},
		},
	}

	for _, tc := range tcs {
		assert.MapEquals(t, tc.rcs, mp.RuneCount(tc.input))
	}
}
