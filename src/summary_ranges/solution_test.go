package summary_ranges

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"strconv"
	"testing"
)

/*
You are given a sorted unique integer array nums.
A range [a,b] is the set of all integers from a to b (inclusive).
Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums
is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
Each range [a,b] in the list should be output as:
 - "a->b" if a != b
 - "a" if a == b

Example 1:
	Input: nums = [0,1,2,4,5,7]
	Output: ["0->2","4->5","7"]
Explanation: The ranges are:
	[0,2] --> "0->2"
	[4,5] --> "4->5"
	[7,7] --> "7"

Example 2:
	Input: nums = [0,2,3,4,6,8,9]
	Output: ["0","2->4","6","8->9"]
Explanation: The ranges are:
	[0,0] --> "0"
	[2,4] --> "2->4"
	[6,6] --> "6"
	[8,9] --> "8->9"

Constraints:
 - 0 <= nums.length <= 20
 - -231 <= nums[i] <= 231 - 1
 - All the values of nums are unique.
 - nums is sorted in ascending order.
*/

func summaryRanges(nums []int) []string {
	var result []string
	if len(nums) == 0 {
		return result
	}
	var start, end int
	var str string
	for i := range nums {
		if i == 0 {
			continue
		}
		if nums[i]-nums[i-1] == 1 {
			end = i
		} else {
			if start == end {
				str = fmt.Sprintf("%d", nums[start])
			} else {
				str = fmt.Sprintf("%d->%d", nums[start], nums[end])
			}
			result = append(result, str)
			start = i
			end = i
		}
	}
	if start == end {
		str = fmt.Sprintf("%d", nums[start])
	} else {
		str = fmt.Sprintf("%d->%d", nums[start], nums[end])
	}
	result = append(result, str)
	return result
}

var testcase = []domains.Testcase{
	{
		In:  []int{0, 1, 2, 4, 5, 7},
		Out: []string{"0->2", "4->5", "7"},
	},
	{
		In:  []int{0, 2, 3, 4, 6, 8, 9},
		Out: []string{"0", "2->4", "6", "8->9"},
	},
}

func summaryRanges2(nums []int) []string {
	var result []string
	r := ""

	for i := 0; i < len(nums); {
		r = r + strconv.Itoa(nums[i])
		j := i + 1
		for ; j < len(nums) && nums[j]-nums[j-1] == 1; j++ {
		}
		if j-1 > i {
			r = r + "->" + strconv.Itoa(nums[j-1])
		}
		result = append(result, r)
		r = ""
		i = j
	}
	return result
}

func TestSummaryRanges(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.([]string)
		assert.Equal(t, output, summaryRanges(input))
		assert.Equal(t, output, summaryRanges2(input))
	}
}
