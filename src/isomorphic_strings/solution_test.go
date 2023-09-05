package isomorphic_strings

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given two strings s and t, determine if they are isomorphic.
Two strings s and t are isomorphic if the characters in s can be replaced to get t.
All occurrences of a character must be replaced with another character while preserving the order of characters. No two
characters may map to the same character, but a character may map to itself.

Example 1:
	Input: s = "egg", t = "add"
	Output: true

Example 2:
	Input: s = "foo", t = "bar"
	Output: false

Example 3:
	Input: s = "paper", t = "title"
	Output: true

Constraints:
 - 1 <= s.length <= 5 * 104
 - t.length == s.length
 - s and t consist of any valid ascii character.
*/

func isIsomorphic(s string, t string) bool {
	_len := len(s)
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < _len; i++ {
		sChar := s[i]
		tChar := t[i]
		if val, ok := sMap[sChar]; ok {
			if val != tChar {
				return false
			}
		} else {
			if _, ok = tMap[tChar]; ok {
				return false
			}
			sMap[sChar] = tChar
			tMap[tChar] = sChar
		}
	}
	return true
}

type Isomorphic struct {
	S string
	T string
}

var testcase = []domains.Testcase{
	{
		In:  Isomorphic{S: "egg", T: "add"},
		Out: true,
	},
	{
		In:  Isomorphic{S: "foo", T: "bar"},
		Out: false,
	},
	{
		In:  Isomorphic{S: "paper", T: "title"},
		Out: true,
	},
	{
		In:  Isomorphic{S: "bbbaaaba", T: "aaabbbba"},
		Out: false,
	},
}

func TestIsomorphicStrings(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(Isomorphic)
		output := tt.Out.(bool)
		assert.Equal(t, output, isIsomorphic(input.S, input.T))
	}
}
