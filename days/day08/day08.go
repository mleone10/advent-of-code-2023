package day08

import (
	"strings"
	"unicode"

	"github.com/mleone10/advent-of-code-2023/internal/slice"
)

type Map struct {
	Input string

	path string
	m    map[string]node
}

type node struct {
	val, left, right string
}

func (m Map) ShortestTraversalDist() int {
	if m.m == nil {
		m.init()
	}

	cur := "AAA"
	dist := 0
	for cur != "ZZZ" {
		d := m.path[dist%len(m.path)]
		// TODO: instead of storing nodes as struct, store them as map[rune]string so that this lookup can be shrunk
		if d == 'L' {
			cur = m.m[cur].left
		} else if d == 'R' {
			cur = m.m[cur].right
		}
		dist++
	}

	return dist
}

func (m *Map) init() {
	iParts := strings.Split(m.Input, "\n\n")
	m.path = iParts[0]

	m.m = slice.Reduce(slice.TrimSplit(iParts[1]), map[string]node{}, func(l string, ns map[string]node) map[string]node {
		nParts := strings.FieldsFunc(l, func(r rune) bool {
			return unicode.IsSpace(r) || unicode.IsSymbol(r) || unicode.IsPunct(r)
		})
		ns[nParts[0]] = node{
			val:   nParts[0],
			left:  nParts[1],
			right: nParts[2],
		}
		return ns
	})
}
