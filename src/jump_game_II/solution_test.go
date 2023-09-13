package jump_game_II

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i],
you can jump to any nums[i + j] where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].

Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to
the last index.

Example 2:

Input: nums = [2,3,0,1,4]
Output: 2


Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 1000
It's guaranteed that you can reach nums[n - 1].
*/

func jump(nums []int) int {
	curJump, farthestJump, jumps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthestJump {
			farthestJump = i + nums[i]
		}

		if i == curJump {
			jumps, curJump = jumps+1, farthestJump
			if curJump >= len(nums)-1 {
				return jumps
			}
		}
	}
	return 0
}

var testcase = []domains.Testcase{
	{
		In:  []int{1, 1, 1, 1},
		Out: 3,
	},
	{
		In:  []int{0},
		Out: 0,
	},
	{
		In:  []int{2, 3, 1, 1, 4},
		Out: 2,
	},
	{
		In:  []int{2, 3, 0, 1, 4},
		Out: 2,
	},
}

func TestJumpGameII(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, jump(input))
	}
}
