package container_with_most_water

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the
ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water
(blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1

Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

func maxArea(height []int) int {
	_maxArea, i, j := 0, 0, len(height)-1
	for i < j {
		if height[i] > height[j] {
			if _maxArea < height[j]*(j-i) {
				_maxArea = height[j] * (j - i)
			}
			j--
		} else {
			if _maxArea < height[i]*(j-i) {
				_maxArea = height[i] * (j - i)
			}
			i++
		}
	}
	return _maxArea
}

var testcases = []domains.Testcase{
	{
		In:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		Out: 49,
	},
	{
		In:  []int{1, 1},
		Out: 1,
	},
}

func TestContainerWithMostWater(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, maxArea(input))
	}
}
