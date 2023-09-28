package hamming_distance

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"math/bits"
	"testing"
)

/*
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
Given two integers x and y, return the Hamming distance between them.



Example 1:

Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
The above arrows point to positions where the corresponding bits are different.
Example 2:

Input: x = 3, y = 1
Output: 1


Constraints:

0 <= x, y <= 2^31 - 1
*/

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func hammingDistance2(x int, y int) int {
	xor, r := x^y, 0

	for xor != 0 {
		xor -= xor & (-xor)
		r++
	}
	return r
}

type HammingDistance struct {
	x int
	y int
}

var testcases = []domains.Testcase{
	{
		In: HammingDistance{
			x: 1,
			y: 4,
		},
		Out: 2,
	},
	{
		In: HammingDistance{
			x: 3,
			y: 1,
		},
		Out: 1,
	},
}

func TestHammingDistance(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(HammingDistance)
		output := tt.Out.(int)
		assert.Equal(t, output, hammingDistance(input.x, input.y))
		assert.Equal(t, output, hammingDistance2(input.x, input.y))
	}
}
