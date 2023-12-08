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

func TestMinMax(t *testing.T) {
	tcs := []struct {
		ns       []int
		min, max int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			1, 5,
		},
	}

	for _, tc := range tcs {
		assert.Equals(t, tc.min, mth.Min(tc.ns...))
		assert.Equals(t, tc.max, mth.Max(tc.ns...))
	}
}

func TestLcm(t *testing.T) {
	tcs := []struct {
		ns       []int
		expected int
	}{
		{
			[]int{10}, 10,
		},
		{
			[]int{10, 15}, 30,
		},
		{
			[]int{10, 15, 20}, 60,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2520,
		},
	}

	for _, tc := range tcs {
		assert.Equals(t, tc.expected, mth.Lcm(tc.ns...))
	}
}
