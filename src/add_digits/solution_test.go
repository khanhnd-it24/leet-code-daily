package add_digits

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.

Example 1:

Input: num = 38
Output: 2
Explanation: The process is
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
Since 2 has only one digit, return it.
Example 2:

Input: num = 0
Output: 0

Constraints:

0 <= num <= 2^31 - 1
*/

// Solution 1:
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	nextNum := 0
	for num > 0 {
		nextNum += num % 10
		num /= 10
	}
	return addDigits(nextNum)
}

// Solution 2:
func addDigits2(num int) int {
	if num == 0 {
		return 0
	}

	return 1 + (num-1)%9
}

var testcase = []domains.Testcase{
	{
		In:  38,
		Out: 2,
	},
	{
		In:  0,
		Out: 0,
	},
}

func TestAddDigits(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(int)
		output := tt.Out.(int)
		assert.Equal(t, output, addDigits(input))
		assert.Equal(t, output, addDigits2(input))
	}
}
