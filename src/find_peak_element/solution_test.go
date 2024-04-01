package find_peak_element

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcases = []domains.Testcase{
	{In: []int{1, 2, 3, 1}, Out: 2},
	{In: []int{1, 2, 1, 3, 5, 6, 4}, Out: 1},
}

func TestFindPeakElement(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, findPeakElement(input))
	}
}
