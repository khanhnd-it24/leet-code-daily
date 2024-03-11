package palindromic_string_in_array

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcases = []domains.Testcase{
	{
		In:  []string{"abc", "car", "ada", "racecar", "cool"},
		Out: "ada",
	},
	{
		In:  []string{"notapalindrome", "racecar"},
		Out: "racecar",
	},
	{
		In:  []string{"def", "ghi"},
		Out: "",
	},
}

func TestPalindromicStringInArray(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([]string)
		output := tt.Out.(string)
		assert.Equal(t, output, firstPalindrome(input))
	}
}
