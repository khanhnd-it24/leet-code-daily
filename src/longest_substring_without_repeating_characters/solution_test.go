package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	seen := make([]int, 256)
	start := 0
	maxLen := 0

	for i := range s {
		start = max(seen[s[i]], start)
		maxLen = max(i-start+1, maxLen)
		seen[s[i]] = i + 1
	}
	return maxLen
}

var testcase = []domains.Testcase{
	{
		In:  "abcabcbb",
		Out: 3,
	},
	{
		In:  "bbbbb",
		Out: 1,
	},
	{
		In:  "pwwkew",
		Out: 3,
	},
}

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(string)
		output := tt.Out.(int)
		assert.Equal(t, output, lengthOfLongestSubstring(input))
	}
}
