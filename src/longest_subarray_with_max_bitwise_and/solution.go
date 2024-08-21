package longest_subarray_with_max_bitwise_and

// Description: https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/description/
/*
 - 1 <= nums.length <= 10^5
 - 1 <= nums[i] <= 10^6
*/

func longestSubarray(nums []int) int {
	maxValue, maxCount, count := nums[0], 1, 1
	for i := 1; i < len(nums); i++ {
		if maxValue < nums[i] {
			maxValue = nums[i]
			count = 1
			maxCount = 1
		} else if nums[i] == maxValue {
			count++
			maxCount = max(maxCount, count)
		} else {
			count = 0
			continue
		}
	}
	return maxCount
}
