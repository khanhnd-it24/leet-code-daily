package divide_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type DivideTwoNumbersInput struct {
	dividend, divisor int
}

var testcase = []domains.Testcase{
	{
		In: DivideTwoNumbersInput{
			dividend: 10,
			divisor:  3,
		},
		Out: 3,
	},
	{
		In: DivideTwoNumbersInput{
			dividend: 7,
			divisor:  -3,
		},
		Out: -2,
	},
	{
		In: DivideTwoNumbersInput{
			dividend: -2147483648,
			divisor:  1,
		},
		Out: -2147483648,
	},
	{
		In: DivideTwoNumbersInput{
			dividend: -2147483648,
			divisor:  -1,
		},
		Out: 2147483647,
	},
}

func TestDivideTwoNumbers(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(DivideTwoNumbersInput)
		output := tt.Out.(int)
		assert.Equal(t, output, divide(input.dividend, input.divisor))
	}
}
