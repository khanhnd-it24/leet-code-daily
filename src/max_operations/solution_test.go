package max_operations

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type MaxOperationsInput struct {
	nums []int
	k    int
}

var testcases = []domains.Testcase{
	{
		In: MaxOperationsInput{
			nums: []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2},
			k:    3,
		},
		Out: 4,
	},
	{
		In: MaxOperationsInput{
			nums: []int{1, 2, 3, 4},
			k:    5,
		},
		Out: 2,
	},
	{
		In: MaxOperationsInput{
			nums: []int{3, 1, 3, 4, 3},
			k:    6,
		},
		Out: 1,
	},
}

func TestMaxOperations(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(MaxOperationsInput)
		output := tt.Out.(int)
		assert.Equal(t, output, maxOperations(input.nums, input.k))
	}
}
