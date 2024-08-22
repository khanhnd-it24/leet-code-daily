package reconstruct_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReconstructMatrix(t *testing.T) {
	type arg struct {
		upper  int
		lower  int
		colsum []int
	}

	type testcase struct {
		arg  arg
		want [][]int
	}

	testcases := []testcase{
		{
			arg: arg{
				upper:  2,
				lower:  1,
				colsum: []int{1, 1, 1},
			},
			want: [][]int{{1, 1, 0}, {0, 0, 1}},
		},
		{
			arg: arg{
				upper:  2,
				lower:  3,
				colsum: []int{2, 2, 1, 1},
			},
			want: nil,
		},
		{
			arg: arg{
				upper:  5,
				lower:  5,
				colsum: []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1},
			},
			want: [][]int{{1, 1, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 0, 0, 1, 1, 0, 1}},
		},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.want, reconstructMatrix(tc.arg.upper, tc.arg.lower, tc.arg.colsum))
	}
}
