package keys_and_rooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_KeysAndRooms(t *testing.T) {
	type testcase struct {
		arg  [][]int
		want bool
	}
	testcases := []testcase{
		{arg: [][]int{{1}, {2}, {3}, {}}, want: true},
		{arg: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, want: false},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.want, canVisitAllRooms(tc.arg))
	}
}
