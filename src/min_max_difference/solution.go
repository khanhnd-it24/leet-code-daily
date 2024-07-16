package min_max_difference

// Description: https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/
func minMaxDifference(num int) int {
	nums := splitNumber(num)
	maxNum := findMax(nums)
	minNum := findMin(nums)
	return maxNum - minNum
}

func findMax(nums []int) int {
	numsMax := make([]int, len(nums))
	copy(numsMax, nums)
	tmp := -1
	for i := len(numsMax) - 1; i >= 0; i-- {
		if numsMax[i] < 9 && tmp == -1 {
			tmp = numsMax[i]
			numsMax[i] = 9
		}
		if tmp != -1 && numsMax[i] == tmp {
			numsMax[i] = 9
		}
	}
	return buildNumber(numsMax)
}

func findMin(nums []int) int {
	numsMin := make([]int, len(nums))
	copy(numsMin, nums)
	tmp := -1
	for i := len(numsMin) - 1; i >= 0; i-- {
		if numsMin[i] > 0 && tmp == -1 {
			tmp = numsMin[i]
			numsMin[i] = 0
		}
		if tmp != -1 && numsMin[i] == tmp {
			numsMin[i] = 0
		}
	}
	return buildNumber(numsMin)
}

func splitNumber(num int) []int {
	out := make([]int, 0)
	for num > 0 {
		out = append(out, num%10)
		num /= 10
	}
	return out
}

func buildNumber(nums []int) int {
	var out int
	for i := len(nums) - 1; i >= 0; i-- {
		out = out*10 + nums[i]
	}
	return out
}
