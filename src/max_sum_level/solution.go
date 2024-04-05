package max_sum_level

import "math"

/*
Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

Example 1:
Input: root = [1,7,0,7,-8,null,null]
Output: 2
Explanation:
Level 1 sum = 1.
Level 2 sum = 7 + 0 = 7.
Level 3 sum = 7 + -8 = -1.
So we return the level with the maximum sum which is level 2.

Example 2:
Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
Output: 2

Constraints:
 - The number of nodes in the tree is in the range [1, 10^4].
 - -10^5 <= Node.val <= 10^5
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	sum := make([]int, 0)
	dfs(root, &sum, 0)
	maxSumInLevel := 0
	for i := 1; i < len(sum); i++ {
		if sum[maxSumInLevel] < sum[i] {
			maxSumInLevel = i
		}
	}
	return maxSumInLevel + 1
}

func dfs(root *TreeNode, sum *[]int, level int) {
	if root == nil {
		return
	}
	if len(*sum) <= level {
		*sum = append(*sum, root.Val)
	} else {
		(*sum)[level] += root.Val
	}
	dfs(root.Left, sum, level+1)
	dfs(root.Right, sum, level+1)
}

func bfs(root *TreeNode) int {
	res := math.MinInt64
	idx := -1
	level := 0

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum := 0
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
		if res < sum {
			res = sum
			idx = level
		}
	}
	return idx
}
