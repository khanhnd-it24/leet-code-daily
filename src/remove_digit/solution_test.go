package remove_digit

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

type RemoveDigitInput struct {
	number string
	digit  byte
}

var testcases = []domains.Testcase{
	{
		In: RemoveDigitInput{
			number: "743",
			digit:  '7',
		},
		Out: "43",
	},
	{
		In: RemoveDigitInput{
			number: "123",
			digit:  '3',
		},
		Out: "12",
	},
	{
		In: RemoveDigitInput{
			number: "1231",
			digit:  '1',
		},
		Out: "231",
	},
	{
		In: RemoveDigitInput{
			number: "551",
			digit:  '5',
		},
		Out: "51",
	},
}

func TestRemoveDigit(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(RemoveDigitInput)
		output := tt.Out.(string)
		assert.Equal(t, output, removeDigit(input.number, input.digit))
	}
}
