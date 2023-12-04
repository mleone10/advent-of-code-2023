package mth_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/mth"
)

func TestPow(t *testing.T) {
	tcs := []struct {
		n, exp, expected int
	}{
		{0, 1, 0},
		{5, 0, 1},
		{2, 3, 8},
	}

	for _, tc := range tcs {
		actual := mth.Pow(tc.n, tc.exp)
		assert.Equals(t, tc.expected, actual)
	}
}

func TestAtoi(t *testing.T) {
	tcs := []struct {
		s        string
		expected int
	}{
		{"0", 0},
		{"25", 25},
		{"foobar", 0},
	}

	for _, tc := range tcs {
		actual := mth.Atoi(tc.s)
		assert.Equals(t, tc.expected, actual)
	}
}
