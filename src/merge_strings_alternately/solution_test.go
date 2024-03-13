package merge_strings_alternately

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type MergeStringsAlternatelyInput struct {
	word1, word2 string
}

var testcase = []domains.Testcase{
	{
		In: MergeStringsAlternatelyInput{
			word1: "abc",
			word2: "pqr",
		},
		Out: "apbqcr",
	},
	{
		In: MergeStringsAlternatelyInput{
			word1: "ab",
			word2: "pqrs",
		},
		Out: "apbqrs",
	},
	{
		In: MergeStringsAlternatelyInput{
			word1: "abcd",
			word2: "pq",
		},
		Out: "apbqcd",
	},
}

func TestMergeStringsAlternately(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(MergeStringsAlternatelyInput)
		output := tt.Out.(string)
		assert.Equal(t, output, mergeAlternately(input.word1, input.word2))
	}
}
