package longest_special_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSpecialSubstring(t *testing.T) {
	type testcase struct {
		arg  string
		want int
	}

	var testcases = []testcase{
		{
			arg:  "acc",
			want: -1,
		},
		{
			arg:  "aaaa",
			want: 2,
		},
		{
			arg:  "abcdef",
			want: -1,
		},
		{
			arg:  "abcaba",
			want: 1,
		},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.want, maximumLength(tc.arg))
	}
}
