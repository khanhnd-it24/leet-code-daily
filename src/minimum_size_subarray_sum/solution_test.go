package minimum_size_subarray_sum

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1
Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0


Constraints:

1 <= target <= 10^9
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^4


Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).
*/

func minSubArrayLen(target int, nums []int) int {
	left, sum := 0, 0
	minLen := len(nums) + 1

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen > len(nums) {
		return 0
	}
	return minLen
}

type MinimumSizeSubarraySum struct {
	Target int
	Nums   []int
}

var testcase = []domains.Testcase{
	{
		In: MinimumSizeSubarraySum{
			Target: 213,
			Nums:   []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
		},
		Out: 8,
	},
	{
		In: MinimumSizeSubarraySum{
			Target: 7,
			Nums:   []int{2, 3, 1, 2, 4, 3},
		},
		Out: 2,
	},
	{
		In: MinimumSizeSubarraySum{
			Target: 4,
			Nums:   []int{1, 4, 4},
		},
		Out: 1,
	},
	{
		In: MinimumSizeSubarraySum{
			Target: 11,
			Nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
		},
		Out: 0,
	},
}

func TestMinimumSizeSubarraySum(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(MinimumSizeSubarraySum)
		output := tt.Out.(int)
		assert.Equal(t, output, minSubArrayLen(input.Target, input.Nums))
	}
}
