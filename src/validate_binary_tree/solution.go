package validate_binary_tree

// Description: https://leetcode.com/problems/validate-binary-tree-nodes
/*
Constraints:
 - n == leftChild.length == rightChild.length
 - 1 <= n <= 10^4
 - -1 <= leftChild[i], rightChild[i] <= n - 1
*/

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	root := findRoot(n, leftChild, rightChild)
	if root == -1 {
		return false
	}

	visited := map[int]bool{root: true}
	queue := []int{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		nodes := []int{leftChild[curr], rightChild[curr]}
		for _, node := range nodes {
			if node != -1 {
				if _, ok := visited[node]; ok {
					return false
				}
				visited[node] = true
				queue = append(queue, node)
			}
		}
	}
	return len(visited) == n
}

func findRoot(n int, left, right []int) int {
	seen := make(map[int]bool)

	for i := 0; i < n; i++ {
		seen[left[i]] = true
		seen[right[i]] = true
	}

	for i := 0; i < n; i++ {
		if !seen[i] {
			return i
		}
	}
	return -1
}
