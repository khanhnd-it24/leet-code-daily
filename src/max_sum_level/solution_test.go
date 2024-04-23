package max_sum_level

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSumLevel(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: -6},
		},
		Right: &TreeNode{Val: 0},
	}
	assert.Equal(t, 2, maxLevelSum(root))
}
