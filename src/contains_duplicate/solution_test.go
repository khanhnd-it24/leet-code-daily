package contains_duplicate

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every
element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true


Constraints:

1 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
*/

func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool)
	for _, num := range nums {
		_, ok := mp[num]
		if ok {
			return true
		}
		mp[num] = true
	}
	return false
}

var testcases = []domains.Testcase{
	{
		In:  []int{1, 2, 3, 1},
		Out: true,
	},
	{
		In:  []int{1, 2, 3, 4},
		Out: false,
	},
	{
		In:  []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		Out: true,
	},
}

func TestContainsDuplicate(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([]int)
		output := tt.Out.(bool)
		assert.Equal(t, output, containsDuplicate(input))
	}
}
