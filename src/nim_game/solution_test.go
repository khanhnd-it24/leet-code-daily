package nim_game

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

func canWinNim(n int) bool {
	return 0xfffffffc|n != 0xfffffffc
}

func canWinNim2(n int) bool {
	return n%4 != 0
}

var testcases = []domains.Testcase{
	{
		In:  1,
		Out: true,
	},
	{
		In:  2,
		Out: true,
	},
	{
		In:  4,
		Out: false,
	},
}

func TestNimGame(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(int)
		output := tt.Out.(bool)
		assert.Equal(t, output, canWinNim(input))
		assert.Equal(t, output, canWinNim2(input))
	}
}
