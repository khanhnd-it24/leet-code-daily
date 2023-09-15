package reverse_words_in_string

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"strings"
	"testing"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	sSplit := strings.Split(s, " ")
	result := strings.Builder{}
	for i := len(sSplit) - 1; i >= 0; i-- {
		if sSplit[i] != "" {
			if i != 0 {
				result.WriteString(sSplit[i] + " ")
			} else {
				result.WriteString(sSplit[i])
			}

		}
	}
	return result.String()
}

func reverseWords2(s string) string {
	words := strings.Fields(s)
	l, h := 0, len(words)-1

	for l < h {
		words[l], words[h] = words[h], words[l]
		l++
		h--
	}

	return strings.TrimSpace(strings.Join(words, " "))
}

var testcases = []domains.Testcase{
	{
		In:  "the sky is blue",
		Out: "blue is sky the",
	},
	{
		In:  "  hello world  ",
		Out: "world hello",
	},
	{
		In:  "a good   example",
		Out: "example good a",
	},
}

func TestHIndex(t *testing.T) {
	for _, tt := range testcases {
		input := tt.In.(string)
		output := tt.Out.(string)
		assert.Equal(t, output, reverseWords(input))
		assert.Equal(t, output, reverseWords2(input))
	}
}
