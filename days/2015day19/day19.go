package day19

import (
	"strings"

	"github.com/mleone10/advent-of-code-2023/internal/set"
	"github.com/mleone10/advent-of-code-2023/internal/str"
)

type Replacements map[string][]string

func NewReplacements(input string) Replacements {
	r := Replacements{}

	for _, l := range strings.Split(strings.TrimSpace(input), "\n") {
		from, to, _ := strings.Cut(l, " => ")
		if _, ok := r[from]; !ok {
			r[from] = []string{}
		}
		r[from] = append(r[from], to)
	}

	return r
}

func CalibrationSum(init string, repl Replacements) int {
	uniqMols := set.Set[string]{}
	for src, opts := range repl {
		for _, o := range opts {
			uniqMols.Add(str.ReplaceVariants(init, src, o)...)
		}
	}
	return uniqMols.Size()
}
