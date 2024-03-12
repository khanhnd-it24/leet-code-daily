package prefix_of_string

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type PrefixOfStringInput struct {
	words []string
	s     string
}

var testcases = []domains.Testcase{
	{
		In: PrefixOfStringInput{
			words: []string{"e", "s", "mi", "e", "ia", "ibwu", "e", "e", "k", "ci", "rip", "suw", "a", "l"},
			s:     "e",
		},
		Out: 4,
	},
	{
		In: PrefixOfStringInput{
			words: []string{"a", "b", "c", "ab", "bc", "abc"},
			s:     "abc",
		},
		Out: 3,
	},
	{
		In: PrefixOfStringInput{
			words: []string{"a", "a"},
			s:     "aa",
		},
		Out: 2,
	},
}

func TestPrefixesOfString(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(PrefixOfStringInput)
		output := tt.Out.(int)
		assert.Equal(t, output, countPrefixes(input.words, input.s))
	}
}
