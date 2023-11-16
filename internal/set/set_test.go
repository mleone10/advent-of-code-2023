package set_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/set"
)

func TestAddAndSize(t *testing.T) {
	s := set.Set[int]{}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	assert.Equals(t, 3, s.Size())
}

func TestAddMultiple(t *testing.T) {
	s := set.Set[int]{}

	s.Add([]int{1, 2, 3}...)

	assert.Equals(t, 3, s.Size())
}
