package permutation_sequence

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type PermutationSequenceInput struct {
	n, k int
}

var testcase = []domains.Testcase{
	{
		In: PermutationSequenceInput{
			n: 1,
			k: 1,
		},
		Out: "1",
	},
	{
		In: PermutationSequenceInput{
			n: 3,
			k: 2,
		},
		Out: "132",
	},
	{
		In: PermutationSequenceInput{
			n: 2,
			k: 2,
		},
		Out: "21",
	},
	{
		In: PermutationSequenceInput{
			n: 3,
			k: 3,
		},
		Out: "213",
	},
	{
		In: PermutationSequenceInput{
			n: 4,
			k: 9,
		},
		Out: "2314",
	},
	{
		In: PermutationSequenceInput{
			n: 3,
			k: 1,
		},
		Out: "123",
	},
}

func TestPermutationSequence(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(PermutationSequenceInput)
		output := tt.Out.(string)
		assert.Equal(t, output, getPermutation(input.n, input.k))
	}
}
