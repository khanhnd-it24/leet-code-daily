package reverse

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  1534236469,
		Out: 0,
	},
	{
		In:  123,
		Out: 321,
	},
	{
		In:  -123,
		Out: -321,
	},
	{
		In:  120,
		Out: 21,
	},
}

func TestDivisorGame(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(int)
		output := tt.Out.(int)
		assert.Equal(t, output, reverse(input))
		assert.Equal(t, output, reverse2(input))
	}
}
