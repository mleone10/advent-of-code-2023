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

// Gcd computes the greatest common divisor of two integers.
// Credit: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func Gcd(a, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

// Lcm computs the least common multiple of at least two integers.
// Credit: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func Lcm(ns ...int) int {
	if len(ns) == 1 {
		return ns[0]
	}

	res := ns[0] * ns[1] / Gcd(ns[0], ns[1])

	for i := 0; i < len(ns[2:]); i++ {
		res = Lcm(res, ns[2+i])
	}

	return res
}
