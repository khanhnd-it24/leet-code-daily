package find_disappeared_numbers

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

var testcase = []domains.Testcase{
	{In: []int{4, 3, 2, 7, 8, 2, 3, 1}, Out: []int{5, 6}},
}

func TestFindDisappearedNumbers(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]int)
		output := tt.Out.([]int)
		assert.Equal(t, output, findDisappearedNumbers(input))
	}
}
