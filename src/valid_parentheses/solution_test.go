package valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  "()",
		Out: true,
	},
	{
		In:  "()[]{}",
		Out: true,
	},
	{
		In:  "(]",
		Out: false,
	},
}

func TestValidParentheses(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(string)
		output := tt.Out.(bool)
		assert.Equal(t, output, isValid(input))
		assert.Equal(t, output, isValid2(input))
	}
}
