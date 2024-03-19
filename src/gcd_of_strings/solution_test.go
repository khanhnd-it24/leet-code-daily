package gcd_of_strings

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type GCDOfStringsInput struct {
	str1, str2 string
}

var testcase = []domains.Testcase{
	{
		In: GCDOfStringsInput{
			str1: "ABCDEF",
			str2: "ABC",
		},
		Out: "",
	},
	{
		In: GCDOfStringsInput{
			str1: "ABCABC",
			str2: "ABC",
		},
		Out: "ABC",
	},
	{
		In: GCDOfStringsInput{
			str1: "ABABAB",
			str2: "ABAB",
		},
		Out: "AB",
	},
	{
		In: GCDOfStringsInput{
			str1: "LEET",
			str2: "CODE",
		},
		Out: "",
	},
}

func TestGCDOfStrings(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(GCDOfStringsInput)
		output := tt.Out.(string)
		assert.Equal(t, output, gcdOfStrings(input.str1, input.str2))
	}
}
