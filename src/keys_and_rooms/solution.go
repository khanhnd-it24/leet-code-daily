package keys_and_rooms

// Description: https://leetcode.com/problems/keys-and-rooms/?envType=study-plan-v2&envId=leetcode-75

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make(map[int]bool)
	queue := make([]int, 0)
	queue = append(queue, rooms[0]...)
	visited[0] = true
	for len(queue) > 0 {
		if _, ok := visited[queue[0]]; ok {
			queue = queue[1:]
			continue
		}
		visited[queue[0]] = true
		queue = append(queue, rooms[queue[0]]...)
		queue = queue[1:]
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}
