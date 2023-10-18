package island_perimeter

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

func islandPerimeter(grid [][]int) int {
	perimeter, onePerimeter, n, m := 0, 0, len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				onePerimeter = 4
				if i > 0 && grid[i-1][j] == 1 {
					onePerimeter -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					onePerimeter -= 2
				}
				perimeter += onePerimeter
			}
		}
	}
	return perimeter
}

var testcase = []domains.Testcase{
	{
		In: [][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		},
		Out: 16,
	},
	{
		In: [][]int{
			{1},
		},
		Out: 4,
	},
	{
		In: [][]int{
			{1, 0},
		},
		Out: 4,
	},
}

func TestIslandPerimeter(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([][]int)
		output := tt.Out.(int)
		assert.Equal(t, output, islandPerimeter(input))
	}
}
