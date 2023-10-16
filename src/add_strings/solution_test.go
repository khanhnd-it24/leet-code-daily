package add_strings

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"strconv"
	"strings"
	"testing"
)

/*
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.
You must solve the problem without using any built-in library for handling large integers (such as BigInteger).
You must also not convert the inputs to integers directly.

Example 1:
Input: num1 = "11", num2 = "123"
Output: "134"

Example 2:
Input: num1 = "456", num2 = "77"
Output: "533"

Example 3:
Input: num1 = "0", num2 = "0"
Output: "0"

Constraints:

1 <= num1.length, num2.length <= 10^4
num1 and num2 consist of only digits.
num1 and num2 don't have any leading zeros except for the zero itself.
*/

func addStrings(num1 string, num2 string) string {
	var result strings.Builder
	num, carry, i, j := 0, 0, len(num1)-1, len(num2)-1
	for i >= 0 || j >= 0 {
		num = carry
		carry = 0
		if i >= 0 {
			num += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			num += int(num2[j] - '0')
			j--
		}
		if num >= 10 {
			carry = num / 10
			num = num % 10
		}
		result.WriteString(strconv.Itoa(num))
	}
	if carry > 0 {
		result.WriteString(strconv.Itoa(carry))
	}
	return reverse(result.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type AddStrings struct {
	Num1, Num2 string
}

var testcase = []domains.Testcase{
	{
		In: AddStrings{
			Num1: "11",
			Num2: "123",
		},
		Out: "134",
	},
	{
		In: AddStrings{
			Num1: "456",
			Num2: "77",
		},
		Out: "533",
	},
	{
		In: AddStrings{
			Num1: "0",
			Num2: "0",
		},
		Out: "0",
	},
	{
		In: AddStrings{
			Num1: "1",
			Num2: "9",
		},
		Out: "10",
	},
}

func TestAddStrings(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(AddStrings)
		output := tt.Out.(string)
		assert.Equal(t, output, addStrings(input.Num1, input.Num2))
	}
}
