package rotate_image

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		Out: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
	},
	{
		In:  [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
		Out: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
	},
}

func TestRotateImage(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([][]int)
		rotate(input)
		output := tt.Out.([][]int)
		assert.Equal(t, output, input)
	}
}
