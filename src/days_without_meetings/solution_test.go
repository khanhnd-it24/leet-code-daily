package days_without_meetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDaysWithoutMeetings(t *testing.T) {
	type arg struct {
		days     int
		meetings [][]int
	}
	type testcase struct {
		arg  arg
		want int
	}

	testcases := []testcase{
		{
			arg: arg{
				days:     10,
				meetings: [][]int{{5, 7}, {1, 3}, {9, 10}},
			},
			want: 2,
		},
		{
			arg: arg{
				days:     5,
				meetings: [][]int{{2, 4}, {1, 3}},
			},
			want: 1,
		},
		{
			arg: arg{
				days:     6,
				meetings: [][]int{{1, 6}},
			},
			want: 0,
		},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.want, countDays(tc.arg.days, tc.arg.meetings), "%d %v", tc.arg.days, tc.arg.meetings)
	}
}
