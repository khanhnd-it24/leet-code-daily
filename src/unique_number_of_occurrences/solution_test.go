package unique_number_of_occurrences

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{In: []int{1, 2, 2, 1, 1, 3}, Out: true},
	{In: []int{1, 2}, Out: false},
	{In: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, Out: true},
}

func TestUniqueNumberOfOccurrences(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.(bool)
		assert.Equal(t, output, uniqueOccurrences(input))
	}
}
