package day01

import (
	"math"
	"strings"

	"github.com/mleone10/advent-of-code-2023/internal/mp"
)

type DigitMap map[string]int

var (
	Numerics      = DigitMap(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9})
	Alphanumerics = DigitMap(mp.Merge(
		Numerics,
		map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}))
)

func CalibrationSum(ls []string, valid DigitMap) int {
	sum := 0
	for _, l := range ls {
		iFirst := math.MaxInt
		first := 0
		iLast := -1
		last := 0
		for str, digit := range valid {
			if i := strings.Index(l, str); i >= 0 {
				if i < iFirst {
					iFirst = i
					first = digit
				}
			}
			if i := strings.LastIndex(l, str); i >= 0 {
				if i > iLast {
					iLast = i
					last = digit
				}
			}
		}
		sum += first*10 + last
	}

	return sum
}
