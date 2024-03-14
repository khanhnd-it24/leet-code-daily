package reverse_vowels

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcases = []domains.Testcase{
	{
		In:  "hello",
		Out: "holle",
	},
	{
		In:  "leetcode",
		Out: "leotcede",
	},
}

func TestReverseVowels(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(string)
		output := tt.Out.(string)
		assert.Equal(t, output, reverseVowels(input))
	}
}
