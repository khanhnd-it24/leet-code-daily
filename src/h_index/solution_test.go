package h_index

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith
paper, return the researcher's h-index.

According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the
given researcher has published at least h papers that have each been cited at least h times.



Example 1:

Input: citations = [3,0,6,1,5]
Output: 3
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5
citations respectively.
Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations
each, their h-index is 3.

Example 2:

Input: citations = [1,3,1]
Output: 1

Constraints:

n == citations.length
1 <= n <= 5000
0 <= citations[i] <= 1000
*/

func hIndex(citations []int) int {
	lenC := len(citations)
	bucket := make([]int, lenC+1)
	for i := 0; i < lenC; i++ {
		if citations[i] >= lenC {
			bucket[lenC]++
		} else {
			bucket[citations[i]]++
		}
	}
	var cnt int
	for i := lenC; i >= 0; i-- {
		cnt += bucket[i]
		if cnt >= i {
			return i
		}
	}

	return -1
}

var testcases = []domains.Testcase{
	{
		In:  []int{3, 0, 6, 1, 5},
		Out: 3,
	},
	{
		In:  []int{1, 3, 1},
		Out: 1,
	},
}

func TestHIndex(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, hIndex(input))
	}
}
