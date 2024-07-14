package days_without_meetings

import (
	"cmp"
	"slices"
)

// Description: https://leetcode.com/problems/count-days-without-meetings/description

func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(a, b []int) int {
		res := cmp.Compare(a[0], b[0])
		if res == 0 {
			return cmp.Compare(a[1], b[1])
		}
		return res
	})

	groupStart, groupEnd := meetings[0][0], meetings[0][1]
	for i := 1; i < len(meetings); i++ {
		if groupEnd >= meetings[i][0]-1 {
			if groupEnd < meetings[i][1] {
				groupEnd = meetings[i][1]
			}
			continue
		}
		days = days - 1 - (groupEnd - groupStart)
		groupStart = meetings[i][0]
		groupEnd = meetings[i][1]
	}

	return days - 1 - (groupEnd - groupStart)
}
