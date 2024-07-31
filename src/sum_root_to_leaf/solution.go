package sum_root_to_leaf

// Description: https://leetcode.com/problems/sum-root-to-leaf-numbers/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	calc(root, 0, &sum)
	return sum
}

func calc(root *TreeNode, currentVal int, sum *int) {
	currentVal = currentVal*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += currentVal
		return
	}
	if root.Left != nil {
		calc(root.Left, currentVal, sum)
	}
	if root.Right != nil {
		calc(root.Right, currentVal, sum)
	}
}
