package height_checker

import (
	"sort"
)

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)

	var count int
	for i := range heights {
		if expected[i] != heights[i] {
			count++
		}
	}

	return count
}
