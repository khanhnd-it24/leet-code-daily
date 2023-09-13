package two_sum_II

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that
they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2]
where 1 <= index1 < index2 < numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.



Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
Example 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].


Constraints:

2 <= numbers.length <= 3 * 10^4
-1000 <= numbers[i] <= 1000
*/

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}

type TwoSumII struct {
	Numbers []int
	Target  int
}

var testcase = []domains.Testcase{
	{
		In: TwoSumII{
			Numbers: []int{2, 7, 11, 15},
			Target:  9,
		},
		Out: []int{1, 2},
	},
	{
		In: TwoSumII{
			Numbers: []int{2, 3, 4},
			Target:  6,
		},
		Out: []int{1, 3},
	},
	{
		In: TwoSumII{
			Numbers: []int{-1, 0},
			Target:  -1,
		},
		Out: []int{1, 2},
	},
}

func TestTwoSumII(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(TwoSumII)
		output := tt.Out.([]int)
		assert.Equal(t, output, twoSum(input.Numbers, input.Target))
	}
}
