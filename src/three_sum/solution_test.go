package three_sum

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"sort"
	"testing"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.


Constraints:

3 <= nums.length <= 3000
-10^5 <= nums[i] <= 10^5
*/

func threeSum(nums []int) [][]int {
	n := len(nums)

	sort.Ints(nums)

	var result [][]int
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}

var testcase = []domains.Testcase{
	{
		In:  []int{-1, 0, 1, 2, -1, -4},
		Out: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		In:  []int{0, 1, 1},
		Out: [][]int{},
	},
	{
		In:  []int{0, 0, 0},
		Out: [][]int{{0, 0, 0}},
	},
}

func TestThreeSum(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.([][]int)
		assert.Equal(t, output, threeSum(input))
	}
}
