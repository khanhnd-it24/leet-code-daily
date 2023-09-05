package long_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{In: []string{"flower", "flow", "flight"}, Out: "fl"},
	{In: []string{"dog", "racecar", "car"}, Out: ""},
	{In: []string{""}, Out: ""},
}

func TestLongCommonPrefix(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]string)
		output := tt.Out.(string)
		assert.Equal(t, output, longestCommonPrefix(input))
	}
}
