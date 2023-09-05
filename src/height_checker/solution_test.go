package height_checker

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{In: []int{1, 1, 4, 2, 1, 3}, Out: 3},
}

func TestHeightChecker(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, heightChecker(input))
	}
}
