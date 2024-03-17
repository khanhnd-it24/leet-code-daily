package highest_altitude

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  []int{-5, 1, 5, 0, -7},
		Out: 1,
	},
	{
		In:  []int{-4, -3, -2, -1, 4, 3, 2},
		Out: 0,
	},
}

func TestHighestAltitude(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(int)
		assert.Equal(t, output, largestAltitude(input))
	}
}
