package int_to_romain

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{In: 3, Out: "III"},
	{In: 58, Out: "LVIII"},
	{In: 1994, Out: "MCMXCIV"},
}

func TestIntToRomain(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(int)
		output := tt.Out.(string)
		assert.Equal(t, output, intToRoman(input))
	}
}
