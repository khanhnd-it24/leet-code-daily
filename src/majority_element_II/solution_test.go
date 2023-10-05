package majority_element_II

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Example 1:
Input: nums = [3,2,3]
Output: [3]

Example 2:
Input: nums = [1]
Output: [1]

Example 3:
Input: nums = [1,2]
Output: [1,2]

Constraints:

1 <= nums.length <= 5 * 10^4
-10^9 <= nums[i] <= 10^9


Follow up: Could you solve the problem in linear time and in O(1) space?
*/

func majorityElement(nums []int) []int {
	firstNum, secondNum, firstCount, secondCount := 0, 1, 0, 0

	for _, num := range nums {
		if num == firstNum {
			firstCount++
		} else if num == secondNum {
			secondCount++
		} else if firstCount == 0 {
			firstNum = num
			firstCount = 1
		} else if secondCount == 0 {
			secondNum = num
			secondCount++
		} else {
			firstCount--
			secondCount--
		}
	}

	firstCount, secondCount = 0, 0
	for _, num := range nums {
		if num == firstNum {
			firstCount++
		} else if num == secondNum {
			secondCount++
		}
	}

	var result []int
	if firstCount > len(nums)/3 {
		result = append(result, firstNum)
	}
	if secondCount > len(nums)/3 {
		result = append(result, secondNum)
	}

	return result
}

var testcases = []domains.Testcase{
	{
		In:  []int{3, 2, 3},
		Out: []int{3},
	},
	{
		In:  []int{1},
		Out: []int{1},
	},
	{
		In:  []int{1, 2},
		Out: []int{1, 2},
	},
}

func TestMajorityElementII(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([]int)
		output := tt.Out.([]int)
		assert.Equal(t, output, majorityElement(input))
	}
}
