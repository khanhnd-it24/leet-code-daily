package product_array_except_self

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of
nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]


Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.


Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space
for space complexity analysis.)
*/

// Solution 1
func productExceptSelf(nums []int) []int {
	product := 1
	countZero := 0
	for _, num := range nums {
		if num != 0 {
			product *= num
		} else {
			countZero++
		}
	}

	for i := range nums {
		if countZero >= 2 {
			nums[i] = 0
		} else if countZero == 1 {
			if nums[i] == 0 {
				nums[i] = product
			} else {
				nums[i] = 0
			}
		} else {
			nums[i] = product / nums[i]
		}
	}
	return nums
}

func productExceptSelf2(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	res[0], res[size-1] = 1, 1

	for i := 1; i < size; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := size - 2; i >= 0; i-- {
		rightProduct *= nums[i+1]
		res[i] *= rightProduct
	}

	return res
}

var testcase = []domains.Testcase{
	{
		In:  []int{1, 2, 3, 4},
		Out: []int{24, 12, 8, 6},
	},
	{
		In:  []int{-1, 1, 0, -3, 3},
		Out: []int{0, 0, 9, 0, 0},
	},
	{
		In:  []int{-1, 0, 0, -3, 3},
		Out: []int{0, 0, 0, 0, 0},
	},
}

func TestProductArrayExceptSelf(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.([]int)
		assert.Equal(t, output, productExceptSelf2(input))
		assert.Equal(t, output, productExceptSelf(input))
	}
}
