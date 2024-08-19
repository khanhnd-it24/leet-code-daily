package average_waiting_time

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageWaitingTime(t *testing.T) {
	type testcase struct {
		arg  [][]int
		want float64
	}

	testcases := []testcase{
		{
			arg:  [][]int{{1, 2}, {2, 5}, {4, 3}},
			want: 5.00000,
		},
		{
			arg:  [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}},
			want: 3.25000,
		},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.want, averageWaitingTime(tc.arg))
	}
}
