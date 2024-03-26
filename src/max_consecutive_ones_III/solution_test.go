package max_consecutive_ones_III

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type MaxConsecutiveOnesIIIInput struct {
	nums []int
	k    int
}

var testcase = []domains.Testcase{
	{In: MaxConsecutiveOnesIIIInput{nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 3}, Out: 10},
	{In: MaxConsecutiveOnesIIIInput{nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 0}, Out: 4},
	{In: MaxConsecutiveOnesIIIInput{nums: []int{0, 0, 0, 1}, k: 4}, Out: 4},
	{In: MaxConsecutiveOnesIIIInput{nums: []int{0, 0, 1, 1, 1, 0, 0}, k: 0}, Out: 3},
	{In: MaxConsecutiveOnesIIIInput{nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, k: 2}, Out: 6},
}

func TestMaxConsecutiveOnesIII(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(MaxConsecutiveOnesIIIInput)
		output := tt.Out.(int)
		assert.Equal(t, output, longestOnes(input.nums, input.k))
	}
}
