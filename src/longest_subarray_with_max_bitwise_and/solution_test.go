package longest_subarray_with_max_bitwise_and

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarrayWithMaximumBitwiseAND(t *testing.T) {
	type testcase struct {
		arg  []int
		want int
	}

	testcases := []testcase{
		{arg: []int{4, 4, 4, 4, 1, 2, 3, 1, 4, 4, 4, 4, 4, 4, 3, 2}, want: 6},
		{arg: []int{311155, 311155, 311155, 311155, 311155, 311155, 311155, 311155, 201191, 311155}, want: 8},
		{arg: []int{1, 2, 3, 3, 2, 2}, want: 2},
		{arg: []int{1, 2, 3, 4}, want: 1},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.want, longestSubarray(tc.arg))
	}
}
