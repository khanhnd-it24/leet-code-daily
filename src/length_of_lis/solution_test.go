package length_of_lis

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  []int{10, 9, 2, 5, 3, 7, 101, 18},
		Out: 4,
	},
	{
		In:  []int{0, 1, 0, 3, 2, 3},
		Out: 4,
	},
	{
		In:  []int{7, 7, 7, 7, 7, 7, 7},
		Out: 1,
	},
	{
		In:  []int{7, 6, 5, 4, 3, 2, 1},
		Out: 1,
	},
}

func TestLengthOfLIS(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, lengthOfLIS(input))
		assert.Equal(t, output, lengthOfLIS2(input))
	}
}
