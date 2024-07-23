package rotting_oranges

// Description: https://leetcode.com/problems/rotting-oranges

// Solution 1

type Point struct {
	i        int
	j        int
	rottenIn int
}

func orangesRotting(grid [][]int) int {
	queue := NewQueue[Point]()
	count := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue.Enqueue(Point{i: i, j: j, rottenIn: 0})
			}
		}
	}
	for queue.IsNotEmpty() {
		e := queue.Dequeue()
		if e.i > 0 && grid[e.i-1][e.j] == 1 {
			grid[e.i-1][e.j] = 2
			queue.Enqueue(Point{i: e.i - 1, j: e.j, rottenIn: e.rottenIn + 1})
			count = max(count, e.rottenIn+1)
		}
		if e.i < m-1 && grid[e.i+1][e.j] == 1 {
			grid[e.i+1][e.j] = 2
			queue.Enqueue(Point{i: e.i + 1, j: e.j, rottenIn: e.rottenIn + 1})
			count = max(count, e.rottenIn+1)
		}
		if e.j > 0 && grid[e.i][e.j-1] == 1 {
			grid[e.i][e.j-1] = 2
			queue.Enqueue(Point{i: e.i, j: e.j - 1, rottenIn: e.rottenIn + 1})
			count = max(count, e.rottenIn+1)
		}
		if e.j < n-1 && grid[e.i][e.j+1] == 1 {
			grid[e.i][e.j+1] = 2
			queue.Enqueue(Point{i: e.i, j: e.j + 1, rottenIn: e.rottenIn + 1})
			count = max(count, e.rottenIn+1)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return count
}

type Queue[T any] struct {
	queue []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		queue: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(e T) {
	q.queue = append(q.queue, e)
}

func (q *Queue[T]) Dequeue() T {
	ele := q.queue[0]
	q.queue = q.queue[1:]
	return ele
}

func (q *Queue[T]) IsNotEmpty() bool {
	return len(q.queue) != 0
}

// Solution 2

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

const (
	FRESH  = 1
	ROTTEN = 2
)

func orangesRotting2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var rottenOranges [][2]int
	numOfFreshOranges := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case FRESH:
				numOfFreshOranges++
			case ROTTEN:
				rottenOranges = append(rottenOranges, [2]int{i, j})
			}
		}
	}

	if numOfFreshOranges == 0 {
		return 0
	}

	minutes := 0
	for len(rottenOranges) > 0 {
		for _, o := range rottenOranges {
			for _, d := range directions {
				i, j := o[0]+d[0], o[1]+d[1]
				if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == FRESH {
					grid[i][j] = ROTTEN
					rottenOranges = append(rottenOranges, [2]int{i, j})
					numOfFreshOranges--
				}
			}
			rottenOranges = rottenOranges[1:]
		}
		minutes++
		if numOfFreshOranges == 0 {
			return minutes
		}
	}

	return -1
}
