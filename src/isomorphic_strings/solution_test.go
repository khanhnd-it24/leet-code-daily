package isomorphic_strings

import (
	"github.com/stretchr/testify/assert"
	"leet-code/src/core/domains"
	"testing"
)

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
	//{
	//	In:  Isomorphic{S: "egg", T: "add"},
	//	Out: true,
	//},
	//{
	//	In:  Isomorphic{S: "foo", T: "bar"},
	//	Out: false,
	//},
	//{
	//	In:  Isomorphic{S: "paper", T: "title"},
	//	Out: true,
	//},
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
