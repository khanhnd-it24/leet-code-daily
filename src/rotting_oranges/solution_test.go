package rotting_oranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RottingOranges(t *testing.T) {
	type testcase struct {
		arg  [][]int
		want int
	}
	testcases := []testcase{
		{
			arg: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			want: 4,
		},
		{
			arg: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			want: -1,
		},
		{
			arg: [][]int{
				{0, 2},
			},
			want: 0,
		},
	}

	for _, tc := range testcases {
		//assert.Equal(t, tc.want, orangesRotting(tc.arg))
		assert.Equal(t, tc.want, orangesRotting2(tc.arg))
	}
}
