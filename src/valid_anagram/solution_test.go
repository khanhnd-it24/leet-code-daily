package valid_anagram

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the
original letters exactly once.

Example 1:
	Input: s = "anagram", t = "nagaram"
	Output: true

Example 2:
	Input: s = "rat", t = "car"
	Output: false

Constraints:
 - 1 <= s.length, t.length <= 5 * 104
 - s and t consist of lowercase English letters.
*/

func isAnagram(s string, t string) bool {
	mp := make([]int, 26)

	for _, v := range s {
		mp[v-'a']++
	}
	for _, v := range t {
		mp[v-'a']--
	}

	for _, v := range mp {
		if v != 0 {
			return false
		}
	}

	return true
}

type Anagram struct {
	S string
	T string
}

var testcase = []domains.Testcase{
	{
		In:  Anagram{S: "anagram", T: "nagaram"},
		Out: true,
	},
	{
		In:  Anagram{S: "rat", T: "bar"},
		Out: false,
	},
}

func TestValidAnagram(t *testing.T) {
	for _, tt := range testcase {
		input := tt.In.(Anagram)
		output := tt.Out.(bool)
		assert.Equal(t, output, isAnagram(input.S, input.T))
	}
}
