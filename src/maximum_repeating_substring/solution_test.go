package maximum_repeating_substring

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcases = []domains.Testcase{
	{
		In:  []string{"aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"},
		Out: 5,
	},
	{
		In:  []string{"babbbbbaab", "ab"},
		Out: 1,
	},
	{
		In:  []string{"aaa", "a"},
		Out: 3,
	},
	{
		In:  []string{"ababc", "ab"},
		Out: 2,
	},
	{
		In:  []string{"ababc", "ba"},
		Out: 1,
	},
	{
		In:  []string{"ababc", "ac"},
		Out: 0,
	},
}

func TestMaximumRepeatingSubstring(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([]string)
		output := tt.Out.(int)
		assert.Equal(t, output, maxRepeating(input[0], input[1]))
	}
}
