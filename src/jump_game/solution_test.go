package jump_game

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the
array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible
to reach the last index.

Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 105
*/

// Solution 1:
func canJump(nums []int) bool {
	var i int
	var farthest int
	for i < len(nums) {
		if i+nums[i] > farthest {
			farthest = nums[i] + i
		}
		if nums[i] == 0 && i < len(nums)-1 && i == farthest {
			return false
		}
		i++
	}
	return true
}

// Solution 2
func canJump2(nums []int) bool {
	target := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= target {
			target = i
		}
	}

	return target == 0
}

var testcase = []domains.Testcase{
	{
		In:  []int{3, 0, 8, 2, 0, 0, 1},
		Out: true,
	},
	{
		In:  []int{2, 3, 1, 1, 4},
		Out: true,
	},
	{
		In:  []int{3, 2, 1, 0, 4},
		Out: false,
	},
}

func TestJumpGame(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(bool)
		assert.Equal(t, output, canJump(input))
		assert.Equal(t, output, canJump2(input))
	}
}
