package valid_mountain_array

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{In: []int{1, 2, 2, 5, 3, 5}, Out: 2},
}

func TestThirdMax(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, thirdMax(input))
	}
}
