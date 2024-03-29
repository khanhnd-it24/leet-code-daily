package asteroid_collision

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcases = []domains.Testcase{
	{
		In:  []int{-2, -1, 1, 2},
		Out: []int{-2, -1, 1, 2},
	},
	{
		In:  []int{5, 10, -5},
		Out: []int{5, 10},
	},
	{
		In:  []int{8, -8},
		Out: []int{},
	},
	{
		In:  []int{10, 2, -5},
		Out: []int{10},
	},
}

func TestAsteroidCollision(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.([]int)
		output := tt.Out.([]int)
		assert.Equal(t, output, asteroidCollision(input))
	}
}
