package longest_alternating_subarray

//You are given a 0-indexed integer array nums. A subarray s of length m is called alternating if:
//
// - m is greater than 1.
// - s1 = s0 + 1.
// - The 0-indexed subarray s looks like [s0, s1, s0, s1,...,s(m-1) % 2]. In other words, s1 - s0 = 1, s2 - s1 = -1, s3 - s2 = 1, s4 - s3 = -1,
// and so on up to s[m - 1] - s[m - 2] = (-1)m.
//Return the maximum length of all alternating subarrays present in nums or -1 if no such subarray exists.
//A subarray is a contiguous non-empty sequence of elements within an array.

func alternatingSubarray(nums []int) int {
	len := len(nums)
	var max int
	m := 0
	a := 1

	for i := 0; i < len; i++ {
		si := nums[i]
		if i == len-1 {
			break
		}

		if m == 0 {
			m++
		}

		if nums[i+1]-si == a {
			m++
			a *= -1
		} else {
			if max < m {
				max = m
			}
			if m > 1 {
				i--
			}
			a = 1
			m = 0
		}
	}

	if max < m {
		max = m
	}

	if max <= 1 {
		return -1
	}

	return max
}
