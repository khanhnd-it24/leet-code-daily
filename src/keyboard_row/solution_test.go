package keyboard_row

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"strings"
	"testing"
)

func findWords(words []string) []string {
	letters := [26]int{}
	for _, r := range "qwertyuiop" {
		letters[r-'a'] = 1
	}
	for _, r := range "asdfghjkl" {
		letters[r-'a'] = 2
	}
	for _, r := range "zxcvbnm" {
		letters[r-'a'] = 3
	}

	var res []string
	for _, word := range words {
		w := strings.ToLower(word)
		match := true
		for i := 1; i < len(w); i++ {
			if letters[w[i]-'a'] != letters[w[0]-'a'] {
				match = false
			}
		}

		if match {
			res = append(res, word)
		}
	}
	return res
}

var testcase = []domains.Testcase{
	{
		In:  []string{"Hello", "Alaska", "Dad", "Peace"},
		Out: []string{"Alaska", "Dad"},
	},
	{
		In:  []string{"omk"},
		Out: []string{},
	},
	{
		In:  []string{"adsdf", "sfd"},
		Out: []string{"adsdf", "sfd"},
	},
}

func TestAddStrings(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.([]string)
		output := tt.Out.([]string)
		assert.Equal(t, output, findWords(input))
	}
}
