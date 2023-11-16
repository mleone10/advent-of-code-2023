package str

import "strings"

// ReplaceVariants is like strings.Replace, but as applied to all occurrences of `old` in `s`.  The returned slice contains all strings resulting from those operations.
func ReplaceVariants(s, old, new string) []string {
	idxs := []int{}
	cur := 0
	for i := 0; i < strings.Count(s, old); i++ {
		idx := strings.Index(s[cur:], old) + cur
		idxs = append(idxs, idx)
		cur = idx + len(old)
	}

	vs := []string{}
	for _, idx := range idxs {
		vs = append(vs, s[:idx]+strings.Replace(s[idx:], old, new, 1))
	}
	return vs
}
