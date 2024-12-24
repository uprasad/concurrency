package graph

import (
	"sync"
)

func ParallelBFS(graph [][]int, start, end Coord, parallelism int) int {
	if len(graph) == 0 || len(graph[0]) == 0 {
		return -1
	}
	if start.X == end.X && start.Y == end.Y {
		return 0
	}

	var queue, nextQueue []Coord
	visited := make(map[Coord]struct{})

	queue = append(queue, start)
	visited[start] = struct{}{}

	dist := 1
	for len(queue) > 0 {
		in := make(chan Coord, parallelism)
		go func() {
			for _, elem := range queue {
				in <- elem
			}
			close(in)
		}()

		for out := range genNeighbors(graph, in, visited, parallelism) {
			if out.X == end.X && out.Y == end.Y {
				return dist
			}

			nextQueue = append(nextQueue, out)
		}

		dist += 1
		queue = nextQueue
		nextQueue = nil
	}

	return -1
}

func genNeighbors(
	graph [][]int,
	in <-chan Coord,
	visited map[Coord]struct{},
	parallelism int,
) <-chan Coord {
	out := make(chan Coord, parallelism)

	go func() {
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(parallelism)

		for i := 0; i < parallelism; i++ {
			go func() {
				defer wg.Done()
				for elem := range in {
					neighbors := getNeighbors(graph, elem)
					mu.Lock()
					for _, n := range neighbors {
						if _, ok := visited[n]; ok {
							continue
						}

						visited[n] = struct{}{}
						out <- n
					}
					mu.Unlock()
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}
