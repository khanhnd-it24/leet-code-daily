package power_of_three

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"math"
	"testing"
)

/*
Given an integer n, return true if it is a power of three. Otherwise, return false.
An integer n is a power of three, if there exists an integer x such that n == 3^x.

Example 1:

Input: n = 27
Output: true
Explanation: 27 = 3^3

Example 2:

Input: n = 0
Output: false
Explanation: There is no x where 3^x = 0.
Example 3:

Input: n = -1
Output: false
Explanation: There is no x where 3^x = (-1).


Constraints:
-2^31 <= n <= 2^31 - 1

Follow up: Could you solve it without loops/recursion?
*/

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	p := math.Log(float64(n)) / math.Log(3)
	if math.Round(p*(10e10))/(10e10)-math.Round(p) == 0 {
		return true
	}
	return false
}

var testcases = []domains.Testcase{
	{
		In:  19682,
		Out: false,
	},
	{
		In:  27,
		Out: true,
	},
	{
		In:  1,
		Out: true,
	},
	{
		In:  0,
		Out: false,
	},
	{
		In:  -1,
		Out: false,
	},
}

func TestPowerOfThree(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(int)
		output := tt.Out.(bool)
		assert.Equal(t, output, isPowerOfThree(input))
	}
}
