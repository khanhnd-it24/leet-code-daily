package palindrome_number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:
	Input: x = 121
	Output: true
	Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
	Input: x = -121
	Output: false
	Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
	Input: x = 10
	Output: false
	Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:

-2^31 <= x <= 2^31 - 1
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := fmt.Sprintf("%d", x)
	i := 0
	for i <= len(s)-1-i {
		if s[i] != s[len(s)-1-i] {
			return false
		}
		i++
	}
	return true
}

var testcase = []domains.Testcase{
	{
		In:  121,
		Out: true,
	},
	{
		In:  -121,
		Out: false,
	},
	{
		In:  10,
		Out: false,
	},
}

func TestPalindromeNumber(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(int)
		output := tt.Out.(bool)
		assert.Equal(t, output, isPalindrome(input))
	}
}
