package atoi

import (
	"leet-code/src/core/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []domains.Testcase{
	{
		In:  "   -88827   5655  U",
		Out: -88827,
	},
	{
		In:  "42",
		Out: 42,
	},
	{
		In:  "   -042",
		Out: -42,
	},
	{
		In:  "1337c0d3",
		Out: 1337,
	},
	{
		In:  "0-1",
		Out: 0,
	},
	{
		In:  "words and 987",
		Out: 0,
	},
	{
		In:  "-91283472332",
		Out: -2147483648,
	},
	{
		In:  "+-12",
		Out: 0,
	},
	{
		In:  "   +   422",
		Out: 0,
	},
}

func TestAsteroidCollision(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(string)
		output := tt.Out.(int)
		assert.Equal(t, output, myAtoi(input))
	}
}
