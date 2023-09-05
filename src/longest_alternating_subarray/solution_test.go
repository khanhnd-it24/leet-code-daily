package longest_alternating_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testcase = [][]interface{}{
	//{4, []int{2, 3, 4, 3, 4}},
	//{2, []int{4, 5, 6}},
	//{4, []int{31, 32, 31, 32, 33}},
	//{-1, []int{21, 39, 1}},
	{8, []int{1, 2, 1, 2, 1, 2, 1, 2}},
}

func TestLongestAlternatingSubArray(t *testing.T) {
	for _, tt := range testcase {
		output := tt[0]
		input := tt[1].([]int)
		assert.Equal(t, output, alternatingSubarray(input), "Input: [%v], output: [%v]", input, output)
	}
}
