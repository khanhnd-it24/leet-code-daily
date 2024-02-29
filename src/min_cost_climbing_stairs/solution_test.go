package min_cost_climbing_stairs

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  []int{10, 15, 20},
		Out: 15,
	},
	{
		In:  []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		Out: 6,
	},
}

func TestLongestPalindrome(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, minCostClimbingStairs(input))
	}
}
