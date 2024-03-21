package pivot_index

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcases = []domains.Testcase{
	{
		In:  []int{-1, -1, -1, -1, 0, 1},
		Out: 1,
	},
	{
		In:  []int{-1, -1, -1, -1, -1, 0},
		Out: 2,
	},
	{
		In:  []int{1, 7, 3, 6, 5, 6},
		Out: 3,
	},
	{
		In:  []int{1, 2, 3},
		Out: -1,
	},
	{
		In:  []int{2, 1, -1},
		Out: 0,
	},
}

func TestFindPivotIndex(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, pivotIndex(input))
	}
}
