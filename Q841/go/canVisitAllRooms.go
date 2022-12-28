package canVisitAllRooms

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, 1000)

	queue := []int{}
	queue = append(queue, rooms[0]...)

	visitedSum := 1
	visited[0] = true

	for len(queue) != 0 {
		now := queue[0]
		if !visited[now] {
			visited[now] = true
			visitedSum++
			queue = append(queue, rooms[now]...)
		}
		queue = queue[1:]
	}

	return visitedSum == len(rooms)
}
