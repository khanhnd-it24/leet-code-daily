package arranging_coins

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcases = []domains.Testcase{
	{
		In:  3,
		Out: 2,
	},
	{
		In:  5,
		Out: 2,
	},
	{
		In:  8,
		Out: 3,
	},
}

func TestArrangingCoins(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(int)
		output := tt.Out.(int)
		assert.Equal(t, output, arrangeCoins(input))
	}
}
