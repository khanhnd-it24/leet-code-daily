package sum_digits_base_k

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given an integer n (in base 10) and a base k, return the sum of the digits of n after converting n from base 10 to base k.

After converting, each digit should be interpreted as a base 10 number, and the sum should be returned in base 10.

Example 1:

Input: n = 34, k = 6
Output: 9
Explanation: 34 (base 10) expressed in base 6 is 54. 5 + 4 = 9.
Example 2:

Input: n = 10, k = 10
Output: 1
Explanation: n is already in base 10. 1 + 0 = 1.


Constraints:

1 <= n <= 100
2 <= k <= 10
*/

func sumBase(n int, k int) int {
	if n < k {
		return n
	}
	return sumBase(n/k, k) + n%k
}

type SumBase struct {
	N int
	K int
}

var testcase = []domains.Testcase{
	{
		In: SumBase{
			N: 34,
			K: 6,
		},
		Out: 9,
	},
	{
		In: SumBase{
			N: 10,
			K: 10,
		},
		Out: 1,
	},
}

func TestSumBase(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(SumBase)
		output := tt.Out.(int)
		assert.Equal(t, output, sumBase(input.N, input.K))
	}
}
