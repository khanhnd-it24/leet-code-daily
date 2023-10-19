package number_complement

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"math/bits"
	"testing"
)

func findComplement(num int) int {
	mask := findMaskNum(0, num)
	return num ^ mask
}

func findMaskNum(mask, num int) int {
	mask <<= 1
	mask += 1
	if mask >= num {
		return mask
	}
	return findMaskNum(mask, num)
}

func findComplement2(num int) int {
	return (1<<(bits.Len(uint(num))) - 1) ^ num
}

var testcase = []domains.Testcase{
	{
		In:  5,
		Out: 2,
	},
	{
		In:  1,
		Out: 0,
	},
}

func TestNumberComplement(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(int)
		output := tt.Out.(int)
		assert.Equal(t, output, findComplement(input))
		assert.Equal(t, output, findComplement2(input))
	}
}
