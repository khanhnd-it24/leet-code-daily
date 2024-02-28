package longest_palindromic_substring

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  "caaaaa",
		Out: "aaaaa",
	},
	{
		In:  "babad",
		Out: "bab",
	},
	{
		In:  "cbbd",
		Out: "bb",
	},
}

func TestLongestPalindrome(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(string)
		output := tt.Out.(string)
		assert.Equal(t, output, longestPalindrome2(input))
	}
}
