package queue_test

import (
	"concurrency/queue"
	"errors"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBufferedQueue(t *testing.T) {
	tests := []struct {
		name  string
		size  int
		elems []int
	}{
		{
			name: "queue with capacity",
			size: 10,
		},
		{
			name: "queue with no capacity",
			size: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := queue.NewBufferedQueue[int](test.size)
			if _, err := q.Pop(); !errors.Is(err, queue.ErrEmpty) {
				t.Errorf("expected error %v, got %v", queue.ErrEmpty, err)
			}

			elems := make([]int, 0, test.size)
			for i := 0; i < test.size; i++ {
				elem := rand.Int()
				elems = append(elems, elem)
				if err := q.Push(elem); err != nil {
					t.Errorf("q.Push: %v", err)
				}
			}

			if err := q.Push(rand.Int()); !errors.Is(err, queue.ErrFull) {
				t.Errorf("expected error %v, got %v", queue.ErrFull, err)
			}

			gotElems := make([]int, 0, test.size)
			for {
				gotElem, err := q.Pop()
				if errors.Is(err, queue.ErrEmpty) {
					break
				}
				if err != nil {
					t.Errorf("q.Pop: %v", err)
				}
				gotElems = append(gotElems, gotElem)
			}

			if len(gotElems) != test.size {
				t.Errorf("expected %d elements, got %d", len(gotElems), test.size)
			}

			if diff := cmp.Diff(elems, gotElems); diff != "" {
				t.Errorf("cmp.Diff: (-want, +got):\n%s", diff)
			}
		})
	}
}
