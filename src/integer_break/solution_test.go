package integer_break

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"math"
	"testing"
)

/*
Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.
Return the maximum product you can get.

Example 1:

Input: n = 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.
Example 2:

Input: n = 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.


Constraints:

2 <= n <= 58
*/

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	countOf3s := n / 3
	remainder := n % 3
	if remainder == 0 {
		return int(math.Pow(3, float64(countOf3s)))
	}
	if remainder == 1 {
		return int(math.Pow(3, float64(countOf3s-1))) * 4
	}
	return int(math.Pow(3, float64(countOf3s))) * 2
}

var testcase = []domains.Testcase{
	{
		In:  2,
		Out: 1,
	},
	{
		In:  10,
		Out: 36,
	},
}

func TestIntegerBreak(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(int)
		output := tt.Out.(int)
		assert.Equal(t, output, integerBreak(input))
	}
}
