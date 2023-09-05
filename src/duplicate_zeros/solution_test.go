package duplicate_zeros

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  []int{1, 0, 2, 3, 0, 4, 5, 0},
		Out: []int{1, 0, 0, 2, 3, 0, 0, 4},
	},
}

func TestDuplicateZeros(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.([]int)
		duplicateZeros(input)
		assert.Equal(t, output, input)
	}
}
