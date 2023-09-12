package remove_duplicate_II

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	idx := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[idx-2] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}

type Output struct {
	K    int
	Nums []int
}

var testcase = []domains.Testcase{
	{
		In: []int{1, 1, 1, 2, 2, 3},
		Out: Output{
			K:    5,
			Nums: []int{1, 1, 2, 2, 3},
		},
	},
	{
		In: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
		Out: Output{
			K:    7,
			Nums: []int{0, 0, 1, 1, 2, 3, 3},
		},
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(Output)
		k := removeDuplicates(input)
		assert.Equal(t, output.K, k)
		for i := 0; i < k; i++ {
			assert.Equal(t, output.Nums[i], input[i])
		}
	}
}
