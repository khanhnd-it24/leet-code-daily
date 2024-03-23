package leaf_similar_trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueNumberOfOccurrences(t *testing.T) {
	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 2},
	}

	assert.Equal(t, false, leafSimilar(root1, root2))
}
