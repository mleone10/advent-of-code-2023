package slice_test

import (
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

var (
	strLength = func(x string) int { return len(x) }
	intDouble = func(x int) int { return 2 * x }
	isEven    = func(x int) bool { return x%2 == 0 }
	isUpper   = func(x string) bool { return x == strings.ToUpper(x) }
)

func TestMap(t *testing.T) {
	assert.ArrayEquals(t, slice.Map([]string{"a", "b", "c"}, strings.ToUpper), []string{"A", "B", "C"})
	assert.ArrayEquals(t, slice.Map([]string{"fizz", "buzz", "hello"}, strLength), []int{4, 4, 5})
	assert.ArrayEquals(t, slice.Map([]int{2, 3, 4}, intDouble), []int{4, 6, 8})
}

func TestFilter(t *testing.T) {
	assert.ArrayEquals(t, slice.Filter([]int{1, 2, 3, 4, 5}, isEven), []int{2, 4})
	assert.ArrayEquals(t, slice.Filter([]string{"FIZZ", "bUZz", "foO", "BAR"}, isUpper), []string{"FIZZ", "BAR"})
}
