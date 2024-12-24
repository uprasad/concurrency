package graph_test

import (
	"concurrency/graph"
	"testing"
)

func TestBFS(t *testing.T) {
	N := 3_000
	hugeGraph := make([][]int, N)
	for i := range hugeGraph {
		hugeGraph[i] = make([]int, N)
		for j := range hugeGraph[i] {
			hugeGraph[i][j] = 1
		}
	}

	tests := []struct {
		name  string
		graph [][]int
		start graph.Coord
		end   graph.Coord

		wantDist int
	}{
		{
			name:  "empty graph",
			graph: [][]int{},
			start: graph.Coord{0, 0},
			end:   graph.Coord{0, 0},

			wantDist: -1,
		},
		{
			name: "same start and end",
			graph: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			start: graph.Coord{1, 1},
			end:   graph.Coord{1, 1},

			wantDist: 0,
		},
		{
			name: "no obstacles",
			graph: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			start: graph.Coord{1, 0},
			end:   graph.Coord{4, 2},

			wantDist: 5,
		},
		{
			name: "some obstacles",
			graph: [][]int{
				{1, 1, 1, 1},
				{1, 0, 0, 1},
				{0, 1, 1, 1},
				{1, 1, 1, 0},
				{1, 1, 1, 1},
			},
			start: graph.Coord{1, 0},
			end:   graph.Coord{4, 2},

			wantDist: 9,
		},
		{
			name: "no path",
			graph: [][]int{
				{1, 1, 1, 1},
				{1, 0, 0, 1},
				{0, 1, 0, 1},
				{1, 1, 1, 0},
				{1, 1, 1, 1},
			},
			start: graph.Coord{1, 0},
			end:   graph.Coord{4, 2},

			wantDist: -1,
		},
		{
			name:  "huge graph",
			graph: hugeGraph,
			start: graph.Coord{0, 0},
			end:   graph.Coord{N - 1, N - 1},

			wantDist: 2*N - 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotDist := graph.BFS(test.graph, test.start, test.end)
			if test.wantDist != gotDist {
				t.Errorf("want distance %d, got %d", test.wantDist, gotDist)
			}
		})
	}
}

func TestParallelBFS(t *testing.T) {
	N := 3_000
	hugeGraph := make([][]int, N)
	for i := range hugeGraph {
		hugeGraph[i] = make([]int, N)
		for j := range hugeGraph[i] {
			hugeGraph[i][j] = 1
		}
	}

	tests := []struct {
		name  string
		graph [][]int
		start graph.Coord
		end   graph.Coord

		wantDist int
	}{
		{
			name:  "empty graph",
			graph: [][]int{},
			start: graph.Coord{0, 0},
			end:   graph.Coord{0, 0},

			wantDist: -1,
		},
		{
			name: "same start and end",
			graph: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			start: graph.Coord{1, 1},
			end:   graph.Coord{1, 1},

			wantDist: 0,
		},
		{
			name: "no obstacles",
			graph: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			start: graph.Coord{1, 0},
			end:   graph.Coord{4, 2},

			wantDist: 5,
		},
		{
			name: "some obstacles",
			graph: [][]int{
				{1, 1, 1, 1},
				{1, 0, 0, 1},
				{0, 1, 1, 1},
				{1, 1, 1, 0},
				{1, 1, 1, 1},
			},
			start: graph.Coord{1, 0},
			end:   graph.Coord{4, 2},

			wantDist: 9,
		},
		{
			name: "no path",
			graph: [][]int{
				{1, 1, 1, 1},
				{1, 0, 0, 1},
				{0, 1, 0, 1},
				{1, 1, 1, 0},
				{1, 1, 1, 1},
			},
			start: graph.Coord{1, 0},
			end:   graph.Coord{4, 2},

			wantDist: -1,
		},
		{
			name:  "huge graph",
			graph: hugeGraph,
			start: graph.Coord{0, 0},
			end:   graph.Coord{N - 1, N - 1},

			wantDist: 2*N - 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotDist := graph.ParallelBFS(test.graph, test.start, test.end, 1)
			if test.wantDist != gotDist {
				t.Errorf("want distance %d, got %d", test.wantDist, gotDist)
			}
		})
	}
}
