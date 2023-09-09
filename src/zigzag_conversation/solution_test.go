package zigzag_conversation

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"strings"
	"testing"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display
this pattern in a fixed font for better legibility)
	P   A   H   N
	A P L S I I G
	Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:
string convert(string s, int numRows);

Example 1:
	Input: s = "PAYPALISHIRING", numRows = 3
	Output: "PAHNAPLSIIGYIR"

Example 2:
	Input: s = "PAYPALISHIRING", numRows = 4
	Output: "PINALSIGYAHRPI"
Explanation:
	P     I    N
	A   L S  I G
	Y A   H R
	P     I
Example 3:

	Input: s = "A", numRows = 1
	Output: "A"

Constraints:
 - 1 <= s.length <= 1000
 - s consists of English letters (lower-case and upper-case), ',' and '.'.
 - 1 <= numRows <= 1000
*/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var result strings.Builder
	lenS := len(s)
	for i := 0; i < numRows; i++ {
		j := 0
		for 2*(numRows-1)*j+i < lenS || 2*(numRows-1)*j-i < lenS {
			if (i > 0 && i < numRows-1) &&
				(2*(numRows-1)*j-i > 0 && 2*(numRows-1)*j-i < lenS) {
				result.WriteByte(s[2*(numRows-1)*j-i])
			}
			if 2*(numRows-1)*j+i < lenS {
				result.WriteByte(s[2*(numRows-1)*j+i])
			}

			j++
		}
	}
	return result.String()
}

type Zigzag struct {
	S       string
	NumRows int
}

var testcase = []domains.Testcase{
	{
		In: Zigzag{
			S:       "PAYPALISHIRING",
			NumRows: 3,
		},
		Out: "PAHNAPLSIIGYIR",
	},
	{
		In: Zigzag{
			S:       "PAYPALISHIRING",
			NumRows: 4,
		},
		Out: "PINALSIGYAHRPI",
	},
	{
		In: Zigzag{
			S:       "A",
			NumRows: 1,
		},
		Out: "A",
	},
	{
		In: Zigzag{
			S:       "AB",
			NumRows: 1,
		},
		Out: "AB",
	},
}

func TestZigzagConversation(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(Zigzag)
		output := tt.Out.(string)
		assert.Equal(t, output, convert(input.S, input.NumRows))
	}
}
