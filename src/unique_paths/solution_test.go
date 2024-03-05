package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type UniquePathsInput struct {
	m, n int
}

var testcase = []domains.Testcase{
	{
		In: UniquePathsInput{
			m: 23,
			n: 12,
		},
		Out: 193536720,
	},
	{
		In: UniquePathsInput{
			m: 3,
			n: 2,
		},
		Out: 3,
	},
	{
		In: UniquePathsInput{
			m: 3,
			n: 7,
		},
		Out: 28,
	},
}

func TestUniquePaths(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(UniquePathsInput)
		output := tt.Out.(int)
		assert.Equal(t, output, uniquePaths(input.m, input.n))
	}
}
