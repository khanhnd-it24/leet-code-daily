package difference_of_two_arrays

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcases = []domains.Testcase{
	{
		In: [][]int{
			{1, 2, 3},
			{2, 4, 6},
		},
		Out: [][]int{
			{1, 3},
			{4, 6},
		},
	},
	{
		In: [][]int{
			{1, 2, 3, 3},
			{1, 1, 2, 2},
		},
		Out: [][]int{
			{3},
			{},
		},
	},
}

func TestFindDifferenceOfTwoArrays(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([][]int)
		output := tt.Out.([][]int)
		assert.EqualValues(t, output, findDifference(input[0], input[1]))
	}
}
