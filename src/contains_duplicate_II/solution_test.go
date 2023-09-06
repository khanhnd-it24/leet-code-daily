package contains_duplicate_II

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such
that nums[i] == nums[j] and abs(i - j) <= k.

Example 1:
	Input: nums = [1,2,3,1], k = 3
	Output: true

Example 2:
	Input: nums = [1,0,1,1], k = 1
	Output: true

Example 3:
	Input: nums = [1,2,3,1,2,3], k = 2
	Output: false

Constraints:
 - 1 <= nums.length <= 105
 - -109 <= nums[i] <= 109
 - 0 <= k <= 105
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)

	for i, v := range nums {
		if val, ok := mp[v]; ok {
			if i-val <= k {
				return true
			}
		}
		mp[v] = i
	}
	return false
}

type ContainsDuplicateII struct {
	Nums []int
	K    int
}

var testcase = []domains.Testcase{
	{
		In:  ContainsDuplicateII{Nums: []int{1, 2, 3, 1}, K: 3},
		Out: true,
	},
	{
		In:  ContainsDuplicateII{Nums: []int{1, 0, 1, 1}, K: 1},
		Out: true,
	},
	{
		In:  ContainsDuplicateII{Nums: []int{1, 2, 3, 1, 2, 3}, K: 2},
		Out: false,
	},
}

func TestContainsDuplicateII(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(ContainsDuplicateII)
		output := tt.Out.(bool)
		assert.Equal(t, output, containsNearbyDuplicate(input.Nums, input.K))
	}
}
