package good_nodes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoodNodes(t *testing.T) {
	input := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 5},
		},
	}

	assert.Equal(t, 4, goodNodes(input))
}
