package queue_test

import (
	"concurrency/queue"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	parallelisms = []int{1, 2, 4, 8, 16}
)

func TestBasicQueue(t *testing.T) {
	tests := []struct {
		name  string
		elems []int
	}{
		{
			name:  "empty",
			elems: []int{},
		},
		{
			name:  "some elements",
			elems: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := queue.NewBasicQueue[int]()
			// Push all elements
			for _, elem := range test.elems {
				q.Push(elem)
			}

			gotElems := make([]int, 0, len(test.elems))
			q.ForEach(func(elem int) {
				gotElems = append(gotElems, elem)
			})

			if diff := cmp.Diff(test.elems, gotElems); diff != "" {
				t.Errorf("cmp.Diff: (-want, +got):\n%s", diff)
			}

			if len(test.elems) != q.Len() {
				t.Errorf("expected queue length %d, got %d", len(test.elems), q.Len())
			}

			// Pop all elements one by one
			for _, elem := range test.elems {
				gotElem, err := q.Pop()
				if err != nil {
					t.Errorf("q.Pop: %v", err)
				}
				if elem != gotElem {
					t.Errorf("want elem %v, got %v", elem, gotElem)
				}
			}

			if q.Len() != 0 {
				t.Errorf("expected queue length %d, got %d", len(test.elems), q.Len())
			}

			// Pop an empty queue
			if _, err := q.Pop(); !errors.Is(err, queue.ErrEmpty) {
				t.Errorf("expected %v, got %v", queue.ErrEmpty, err)
			}
		})
	}
}

func BenchmarkLockQueue_10_000(b *testing.B) {
	for _, p := range parallelisms {
		q := queue.NewLockQueue[int]()
		b.Run(fmt.Sprintf("Push-%d", p), func(b *testing.B) {
			benchmarkQueuePush(b, q, p, 10_000)
			// b.Logf("expected count: %d, got %d\n", 10_000*p, q.Len())
		})

		b.Run(fmt.Sprintf("Push-Pop-%d", p), func(b *testing.B) {
			benchmarkQueuePushPop(b, q, p, 1_000)
			// b.Logf("count %d\n", q.Len())
		})
	}
}

func BenchmarkLockQueue_100_000(b *testing.B) {
	for _, p := range parallelisms {
		q := queue.NewLockQueue[int]()
		b.Run(fmt.Sprintf("Push-%d", p), func(b *testing.B) {
			benchmarkQueuePush(b, q, p, 100_000)
			// b.Logf("expected count: %d, got %d\n", 100_000*p, q.Len())
		})

		b.Run(fmt.Sprintf("Push-Pop-%d", p), func(b *testing.B) {
			benchmarkQueuePushPop(b, q, p, 10_000)
			// b.Logf("count %d\n", q.Len())
		})
	}
}

func benchmarkQueuePush(b *testing.B, q queue.Queue[int], parallelism, count int) {
	elems := make([]int, 0, count)
	for i := 0; i < count; i++ {
		elems = append(elems, rand.Int())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Clear()
		// Kick off goroutines that increment in parallel
		var wg sync.WaitGroup
		wg.Add(parallelism)
		for j := 0; j < parallelism; j++ {
			go func() {
				defer wg.Done()
				pushElems(q, elems)
			}()
		}
		wg.Wait()
	}
}

func benchmarkQueuePushPop(b *testing.B, q queue.Queue[int], parallelism, count int) {
	elems := make([]int, 0, count)
	for i := 0; i < count; i++ {
		elems = append(elems, rand.Int())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Kick off goroutines that update the queue
		var wg sync.WaitGroup
		wg.Add(parallelism)
		for j := 0; j < parallelism; j++ {
			go func() {
				defer wg.Done()
				if rand.Int()%2 == 0 {
					pushElems(q, elems)
				} else {
					popElems(q, count)
				}
			}()
		}
		wg.Wait()
	}
}

func pushElems(q queue.Queue[int], elems []int) {
	for _, elem := range elems {
		q.Push(elem)
	}
}

func popElems(q queue.Queue[int], count int) {
	for i := 0; i < count; i++ {
		q.Pop()
	}
}
