package divisor_game

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  1,
		Out: false,
	},
	{
		In:  2,
		Out: true,
	},
	{
		In:  3,
		Out: false,
	},
}

func TestDivisorGame(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(int)
		output := tt.Out.(bool)
		assert.Equal(t, output, divisorGame(input))
	}
}
