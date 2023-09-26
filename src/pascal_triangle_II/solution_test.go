package pascal_triangle_II

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

func getRow(rowIndex int) []int {
	var result []int
	x := 1
	result = append(result, x)
	for i := 0; i < rowIndex; i++ {
		x = x * (rowIndex - i) / (i + 1)
		result = append(result, x)
	}
	return result
}

var testcases = []domains.Testcase{
	{
		In:  3,
		Out: []int{1, 3, 3, 1},
	},
	{
		In:  0,
		Out: []int{1},
	},
	{
		In:  1,
		Out: []int{1, 1},
	},
}

func TestPascalTriangleII(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(int)
		output := tt.Out.([]int)
		assert.Equal(t, output, getRow(input))
	}
}
