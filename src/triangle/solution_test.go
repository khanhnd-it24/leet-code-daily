package triangle

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given a triangle array, return the minimum path sum from top to bottom.
For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current
row, you may move to either index i or index i + 1 on the next row.



Example 1:

Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
Example 2:

Input: triangle = [[-10]]
Output: -10


Constraints:

1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-10^4 <= triangle[i][j] <= 10^4


Follow up: Could you do this using only O(n) extra space, where n is the total number of rows in the triangle?
*/

func minimumTotal(triangle [][]int) int {
	ans := make([][]int, len(triangle))
	for i, v := range triangle {
		ans[i] = make([]int, len(v))
		for j, _ := range v {
			ans[i][j] = -1
		}
	}
	return minT(triangle, 0, 0, ans)
}

func minT(triangle [][]int, i, j int, ans [][]int) int {
	if i >= len(triangle) {
		return 0
	}

	if ans[i][j] != -1 {
		return ans[i][j]
	}

	ans[i][j] = triangle[i][j] + min(minT(triangle, i+1, j, ans), minT(triangle, i+1, j+1, ans))

	return ans[i][j]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

var testcases = []domains.Testcase{
	{
		In:  [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
		Out: 11,
	},
	{
		In:  [][]int{{-10}},
		Out: -10,
	},
}

func TestTriangle(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([][]int)
		output := tt.Out.(int)
		assert.Equal(t, output, minimumTotal(input))
	}
}
