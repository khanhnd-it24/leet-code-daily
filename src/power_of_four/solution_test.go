package power_of_four

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given an integer n, return true if it is a power of four. Otherwise, return false.
An integer n is a power of four, if there exists an integer x such that n == 4^x.

Example 1:

Input: n = 16
Output: true
Example 2:

Input: n = 5
Output: false
Example 3:

Input: n = 1
Output: true


Constraints:

-2^31 <= n <= 2^31 - 1
Follow up: Could you solve it without loops/recursion?
*/

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	return n%4 == 0 && isPowerOfFour(n/4)
}

/*
1. The number is a power of 2 for example : 4, 16, 64 … all are powers of 2 as well . The O(1) bit manipulation technique
using n&(n-1)==0 can be used, but wait every power of 4 is a power of 2. Will the reverse apply?
2. Obviously not, 8,32,128 … aren’t power of 4, so we need one more check here. If you clearly notice all powers of 4 when
divided by 3 gives 1 as the remainder.
*/
func isPowerOfFour2(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && (n%3 == 1)
}

var testcases = []domains.Testcase{
	{
		In:  16,
		Out: true,
	},
	{
		In:  5,
		Out: false,
	},
	{
		In:  1,
		Out: true,
	},
}

func TestPowerOfFour(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(int)
		output := tt.Out.(bool)
		assert.Equal(t, output, isPowerOfFour(input))
		assert.Equal(t, output, isPowerOfFour2(input))
	}
}
