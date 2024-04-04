package longest_zigzag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestZigZag(t *testing.T) {
	input := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:   1,
						Right: &TreeNode{Val: 1},
					},
				},
				Right: &TreeNode{Val: 1},
			},
		},
	}

	assert.Equal(t, 3, longestZigZag(input))
}
