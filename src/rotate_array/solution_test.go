package rotate_array

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

type RotateArrayInput struct {
	Nums []int
	K    int
}

var testcase = []domains.Testcase{
	{
		In: RotateArrayInput{
			Nums: []int{1, 2, 3, 4, 5, 6, 7},
			K:    3,
		},
		Out: []int{5, 6, 7, 1, 2, 3, 4},
	},
}

func TestRotateArray(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(RotateArrayInput)
		output := tt.Out.([]int)
		rotate(input.Nums, input.K)
		assert.Equal(t, output, input.Nums)
	}
}
