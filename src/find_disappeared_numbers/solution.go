package find_disappeared_numbers

/*
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:
	Input: nums = [4,3,2,7,8,2,3,1]
	Output: [5,6]

Example 2:
	Input: nums = [1,1]
	Output: [2]

Constraints:
 - n == nums.length
 - 1 <= n <= 105
 - 1 <= nums[i] <= n
*/

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		if i+1 == nums[i] || nums[nums[i]-1] == nums[i] {
			i++
		} else {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	var result []int
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			result = append(result, i+1)
		}
	}

	return result
}
