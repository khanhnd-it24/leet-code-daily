package house_robber

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{In: []int{1, 2, 3, 1}, Out: 4},
	{In: []int{2, 7, 9, 3, 1}, Out: 12},
}

func TestHouseRobber(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, rob(input))
	}
}
