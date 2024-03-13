package length_of_lis

import "sort"

/*
Given an integer array nums, return the length of the longest strictly increasing subsequence

Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1

Constraints:
 - 1 <= nums.length <= 2500
 - -104 <= nums[i] <= 104
 - Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
*/

// DP
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	maxLength := dp[0]
	for i := 1; i < n; i++ {
		if maxLength < dp[i] {
			maxLength = dp[i]
		}
	}
	return maxLength
}

// DP vs Binary Search
func lengthOfLIS2(nums []int) int {
	lis := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > lis[len(lis)-1] {
			lis = append(lis, nums[i])
		} else {
			j := sort.Search(len(lis), func(m int) bool {
				return lis[m] >= nums[i]
			})
			lis[j] = nums[i]
		}
	}
	return len(lis)
}
