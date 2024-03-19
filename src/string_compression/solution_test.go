package string_compression

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{
		In:  []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
		Out: 6,
	},
	{
		In:  []byte{'a'},
		Out: 1,
	},
	{
		In:  []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
		Out: 4,
	},
}

func TestStringCompression(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]byte)
		output := tt.Out.(int)
		assert.Equal(t, output, compress(input))
	}
}
