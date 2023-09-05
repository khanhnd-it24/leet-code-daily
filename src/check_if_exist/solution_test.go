package check_if_exist

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  []int{10, 2, 5, 3},
		Out: true,
	},
	{
		In:  []int{3, 1, 7, 11},
		Out: false,
	},
}

func TestCheckIfExist(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(bool)
		assert.Equal(t, output, checkIfExist(input))
	}
}
