package graph

func BFS(graph [][]int, start, end Coord) int {
	if len(graph) == 0 || len(graph[0]) == 0 {
		return -1
	}

	var queue, nextQueue []Coord
	queue = append(queue, start)

	visited := make(map[Coord]struct{})
	visited[start] = struct{}{}

	dist := 0
	for len(queue) > 0 {
		for _, elem := range queue {
			if elem.X == end.X && elem.Y == end.Y {
				return dist
			}

			neighbors := getNeighbors(graph, elem)

			for _, neighbor := range neighbors {
				if _, ok := visited[neighbor]; ok {
					continue
				}

				visited[neighbor] = struct{}{}
				nextQueue = append(nextQueue, neighbor)
			}
		}

		dist += 1
		queue = nextQueue
		nextQueue = nil
	}

	return -1
}

func getNeighbors(graph [][]int, coord Coord) []Coord {
	M := len(graph)
	N := len(graph[0])

	var neighbors []Coord
	deltas := []Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range deltas {
		n := Coord{X: coord.X + d.X, Y: coord.Y + d.Y}
		if n.X >= 0 && n.X < M && n.Y >= 0 && n.Y < N && graph[n.X][n.Y] != 0 {
			neighbors = append(neighbors, n)
		}
	}

	return neighbors
}
