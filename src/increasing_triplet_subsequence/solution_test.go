package increasing_triplet_subsequence

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  []int{0, 4, 2, 1, 0, -1, -3},
		Out: false,
	},
	{
		In:  []int{1, 2, 3, 4, 5},
		Out: true,
	},
	{
		In:  []int{5, 4, 3, 2, 1},
		Out: false,
	},
	{
		In:  []int{2, 1, 5, 0, 4, 6},
		Out: true,
	},
}

func TestIncreasingTripletSubsequence(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(bool)
		assert.Equal(t, output, increasingTriplet(input))
	}
}
