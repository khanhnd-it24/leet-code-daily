package max_vowels

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type MaxVowelsInput struct {
	s string
	k int
}

var testcase = []domains.Testcase{
	{In: MaxVowelsInput{s: "abciiidef", k: 3}, Out: 3},
	{In: MaxVowelsInput{s: "aeiou", k: 2}, Out: 2},
	{In: MaxVowelsInput{s: "leetcode", k: 3}, Out: 2},
}

func TestMaxVowels(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(MaxVowelsInput)
		output := tt.Out.(int)
		assert.Equal(t, output, maxVowels(input.s, input.k))
	}
}
