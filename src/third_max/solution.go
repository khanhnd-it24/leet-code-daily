package valid_mountain_array

import "math"

func thirdMax(nums []int) int {
	max, max2, max3 := math.MinInt32-1, math.MinInt32-1, math.MinInt32-1

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max3 = max2
			max2 = max
			max = nums[i]
		} else if nums[i] > max2 && nums[i] != max {
			max3 = max2
			max2 = nums[i]
		} else if nums[i] > max3 && nums[i] != max2 && nums[i] != max {
			max3 = nums[i]
		}
	}

	if max3 == math.MinInt32-1 {
		return max
	}
	return max3
}
