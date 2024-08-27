package rank_team

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RankTeams(t *testing.T) {
	type testcase struct {
		arg  []string
		want string
	}

	testcases := []testcase{
		{
			arg:  []string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"},
			want: "ABC",
		},
		{
			arg:  []string{"ABC", "ACB", "ABC", "ACB", "ACB"},
			want: "ACB",
		},
		{
			arg:  []string{"WXYZ", "XYZW"},
			want: "XWYZ",
		},
		{
			arg:  []string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"},
			want: "ZMNAGUEDSJYLBOPHRQICWFXTVK",
		},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.want, rankTeams(tc.arg))
	}
}
