package average_waiting_time

import "math"

// Description: https://leetcode.com/problems/average-waiting-time
/*
	- 1 <= customers.length <= 10^5
	- 1 <= arrival[i], time[i] <= 10^4
	- arrival[i] <= arrival[i+1]
*/
const (
	unit = 100000
)

func averageWaitingTime(customers [][]int) float64 {
	currEnd := 0
	totalWaitTime := 0
	for _, customer := range customers {
		if currEnd < customer[0] {
			currEnd = customer[0] + customer[1]
		} else {
			currEnd += customer[1]
		}
		totalWaitTime += currEnd - customer[0]
	}
	return math.Round(float64(totalWaitTime)/float64(len(customers))*unit) / unit
}
