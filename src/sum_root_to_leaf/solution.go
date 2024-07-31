package sum_root_to_leaf

// Description: https://leetcode.com/problems/sum-root-to-leaf-numbers/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	ans := 0
	compute(root, &ans, 0)
	return ans
}

func compute(root *TreeNode, ans *int, temp int) {
	temp = temp*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*ans += temp
		return
	}
	if root.Left != nil {
		compute(root.Left, ans, temp)
	}
	if root.Right != nil {
		compute(root.Right, ans, temp)
	}
}
