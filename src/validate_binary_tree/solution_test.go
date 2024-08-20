package validate_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateBinaryTree(t *testing.T) {
	type arg struct {
		n     int
		left  []int
		right []int
	}
	type testcase struct {
		arg  arg
		want bool
	}

	testcases := []testcase{
		{
			arg: arg{
				n:     4,
				left:  []int{1, -1, 3, -1},
				right: []int{2, -1, -1, -1},
			},
			want: true,
		},
		{
			arg: arg{
				n:     4,
				left:  []int{1, -1, 3, -1},
				right: []int{2, 3, -1, -1},
			},
			want: false,
		},
		{
			arg: arg{
				n:     2,
				left:  []int{1, 0},
				right: []int{-1, -1},
			},
			want: false,
		},
		{
			arg: arg{
				n:     6,
				left:  []int{1, -1, -1, 4, -1, -1},
				right: []int{2, -1, -1, 5, -1, -1},
			},
			want: false,
		},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.want, validateBinaryTreeNodes(tc.arg.n, tc.arg.left, tc.arg.right))
	}
}
