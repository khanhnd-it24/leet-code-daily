package sum_root_to_leaf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumRootToLeaf(t *testing.T) {
	type testcase struct {
		arg  *TreeNode
		want int
	}

	testcases := []testcase{
		{
			arg: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: 25,
		},
		{
			arg: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   9,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 0},
			},
			want: 1026,
		},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.want, sumNumbers(tc.arg))
	}
}
