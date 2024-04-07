package max_product_difference

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcases = []domains.Testcase{
	{
		In:  []int{5, 6, 2, 7, 4},
		Out: 34,
	},
	{
		In:  []int{4, 2, 5, 9, 7, 4, 8},
		Out: 64,
	},
}

func TestMaxProductDifference(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, maxProductDifference(input))
	}
}
