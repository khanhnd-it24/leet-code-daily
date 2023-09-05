package valid_mountain_array

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	//{In: []int{2, 1}, Out: false},
	//{In: []int{3, 5, 5}, Out: false},
	//{In: []int{0, 3, 2, 1}, Out: true},
	{In: []int{3, 7, 6, 4, 3, 0, 1, 0}, Out: false},
}

func TestValidMountainArray(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(bool)
		assert.Equal(t, output, validMountainArray(input))
	}
}
