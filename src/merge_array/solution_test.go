package merge_array

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type MergeInput struct {
	Nums1 []int
	M     int
	Nums2 []int
	N     int
}

var testcase = []domains.Testcase{
	{
		In: MergeInput{
			Nums1: []int{1, 2, 3, 0, 0, 0},
			M:     3,
			Nums2: []int{2, 5, 6},
			N:     3,
		},
		Out: []int{1, 2, 2, 3, 5, 6},
	},
	{
		In: MergeInput{
			Nums1: []int{1},
			M:     1,
			Nums2: []int{},
			N:     0,
		},
		Out: []int{1},
	},
	{
		In: MergeInput{
			Nums1: []int{2, 0},
			M:     1,
			Nums2: []int{1},
			N:     1,
		},
		Out: []int{1, 2},
	},
}

func TestMergeArray(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(MergeInput)
		output := tt.Out.([]int)
		merge(input.Nums1, input.M, input.Nums2, input.N)
		assert.Equal(t, output, input.Nums1)
	}
}
