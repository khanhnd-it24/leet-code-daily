package make_the_string_great

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeTheStringGreat(t *testing.T) {
	type testcase struct {
		arg  string
		want string
	}
	testcases := []testcase{
		{arg: "leEeetcode", want: "leetcode"},
		{arg: "abBAcC", want: ""},
		{arg: "s", want: "s"},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.want, makeGood(tc.arg))
		assert.Equal(t, tc.want, makeGood2(tc.arg))
	}
}
