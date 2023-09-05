package word_pattern

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type WordPattern struct {
	Pattern string
	S       string
}

var testcases = []domains.Testcase{
	{
		In: WordPattern{
			Pattern: "abba",
			S:       "dog cat cat dog",
		},
		Out: true,
	},
	{
		In: WordPattern{
			Pattern: "abba",
			S:       "dog cat cat fish",
		},
		Out: false,
	},
	{
		In: WordPattern{
			Pattern: "aaaa",
			S:       "dog cat cat dog",
		},
		Out: false,
	},
}

func TestWordPattern(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(WordPattern)
		output := tt.Out.(bool)
		assert.Equal(t, output, wordPattern(input.Pattern, input.S))
	}
}
