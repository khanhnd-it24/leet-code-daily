package find_disappeared_numbers

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
