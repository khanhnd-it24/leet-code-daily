package longest_subarray_ones

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{In: []int{1, 1, 0, 1}, Out: 3},
	{In: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, Out: 5},
	{In: []int{1, 1, 1}, Out: 2},
}

func TestLongestSubarrayOnes(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, longestSubarray(input))
	}
}
