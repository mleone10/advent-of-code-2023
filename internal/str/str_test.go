package str_test

import (
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/assert"
	"github.com/mleone10/advent-of-code-2023/internal/str"
)

func TestReplaceVariants(t *testing.T) {
	tcs := []struct {
		s, old, new string
		expected    []string
	}{
		{
			"AxxBxxCxx", "xx", "yy",
			[]string{"AyyBxxCxx", "AxxByyCxx", "AxxBxxCyy"},
		},
		{
			"AxxBxxCxx", "xx", "zzz",
			[]string{"AzzzBxxCxx", "AxxBzzzCxx", "AxxBxxCzzz"},
		},
		{
			"AxxxBxxCxxx", "xx", "yy",
			[]string{"AyyxBxxCxxx", "AxxxByyCxxx", "AxxxBxxCyyx"},
		},
	}

	for _, tc := range tcs {
		actual := str.ReplaceVariants(tc.s, tc.old, tc.new)

		assert.ArrayEquals(t, tc.expected, actual)
	}

}
