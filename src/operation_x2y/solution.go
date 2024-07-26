package operation_x2y

// Description:https://leetcode.com/problems/minimum-number-of-operations-to-make-x-and-y-equal

type Node struct {
	value int
	level int
}

func minimumOperationsToMakeEqual(x int, y int) int {
	if x <= y {
		return y - x
	}
	queue := []*Node{{x, 0}}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.value == y {
			return node.level
		}

		if _, ok := visited[node.value]; ok {
			continue
		}

		visited[node.value] = true
		if node.value%11 == 0 {
			queue = append(queue, &Node{node.value / 11, node.level + 1})
		}
		if node.value%5 == 0 {
			queue = append(queue, &Node{node.value / 5, node.level + 1})
		}
		queue = append(queue, &Node{node.value - 1, node.level + 1})
		queue = append(queue, &Node{node.value + 1, node.level + 1})
	}
	return -1
}
