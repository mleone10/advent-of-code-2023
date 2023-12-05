package mth

import (
	"math"
	"strconv"
)

// Pow returns the result of `n` to the power `exp`.  It is a convenience method over using math.Pow with type casting.
func Pow(n, exp int) int {
	return int(math.Pow(float64(n), float64(exp)))
}

// Atoi returns the result of `strconv.Atoi`, or 0 if conversion is not possible.
func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// Min returns the smallest integer within `ns`.
func Min(ns ...int) int {
	min := math.MaxInt
	for _, n := range ns {
		if n < min {
			min = n
		}
	}
	return min
}

// Max returns the largest integer within `ns`.
func Max(ns ...int) int {
	max := 0
	for _, n := range ns {
		if n > max {
			max = n
		}
	}
	return max
}
