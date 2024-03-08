package range_addition_II

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type RangeAdditionIIInput struct {
	m, n int
	ops  [][]int
}

var testcases = []domains.Testcase{
	{
		In: RangeAdditionIIInput{
			m:   3,
			n:   3,
			ops: [][]int{{2, 2}, {3, 3}},
		},
		Out: 4,
	},
	{
		In: RangeAdditionIIInput{
			m:   3,
			n:   3,
			ops: [][]int{{2, 2}, {3, 3}},
		},
		Out: 4,
	},
	{
		In: RangeAdditionIIInput{
			m:   3,
			n:   3,
			ops: [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}},
		},
		Out: 4,
	},
	{
		In: RangeAdditionIIInput{
			m:   3,
			n:   3,
			ops: [][]int{},
		},
		Out: 9,
	},
}

func TestRangeAdditionII(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(RangeAdditionIIInput)
		output := tt.Out.(int)
		assert.Equal(t, output, maxCount(input.m, input.n, input.ops))
	}
}
