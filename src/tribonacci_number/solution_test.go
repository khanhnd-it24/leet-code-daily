package tribonacci_number

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  4,
		Out: 4,
	},
	{
		In:  25,
		Out: 1389537,
	},
}

func TestNthTribonacciNumber(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(int)
		output := tt.Out.(int)
		assert.Equal(t, output, tribonacci(input))
	}
}
