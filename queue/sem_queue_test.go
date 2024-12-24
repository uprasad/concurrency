package queue_test

import (
	"concurrency/queue"
	"slices"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSemQueue_PushThenPop(t *testing.T) {
	q := queue.NewSemQueue[int](1_000_000)

	// Push elements 1 to count into the queue
	count := 10_000
	wantElems := make([]int, 0, count)
	for i := 0; i < count; i++ {
		wantElems = append(wantElems, i+1)
	}

	// Insert a bunch of elements concurrently
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := i; j < count; j += 10 {
				if err := q.Push(j + 1); err != nil {
					t.Errorf("q.Push: %v", err)
				}
			}
		}()
	}
	wg.Wait()

	gotElems := make([]int, 0, count)
	q.ForEach(func(elem int) {
		gotElems = append(gotElems, elem)
	})

	slices.Sort(gotElems)

	if diff := cmp.Diff(wantElems, gotElems); diff != "" {
		t.Errorf("cmp.Diff push: (-want, +got)\n%s", diff)
	}

	// Pop all the elements concurrently
	wg.Add(20)
	gotCh := make(chan int, count)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			var got []int
			for j := i; j < count; j += 20 {
				v, err := q.Pop()
				if err != nil {
					t.Errorf("q.Push: %v", err)
				}
				got = append(got, v)
			}

			for _, v := range got {
				gotCh <- v
			}
		}()
	}
	wg.Wait()
	close(gotCh)

	gotPopped := make([]int, 0, count)
	for v := range gotCh {
		gotPopped = append(gotPopped, v)
	}

	slices.Sort(gotPopped)

	if diff := cmp.Diff(wantElems, gotPopped); diff != "" {
		t.Errorf("cmp.Diff pop: (-want, +got)\n%s", diff)
	}
}

func TestSemQueue_PushAndPop(t *testing.T) {
	q := queue.NewSemQueue[int](1_000_000)

	// Push elements 1 to pushCount into the queue
	pushCount := 10_000
	popCount := 1_000
	wantElems := make([]int, 0, pushCount)
	for i := 0; i < pushCount; i++ {
		wantElems = append(wantElems, i+1)
	}

	// Push to and pop from the queue concurrently
	var wg sync.WaitGroup
	wg.Add(10 + 20)
	popCh := make(chan int, popCount)
	for i := 0; i < 10; i++ {
		var pops []int
		go func() {
			defer wg.Done()
			for j := 0; j < popCount/10; j++ {
				v, err := q.Pop()
				if err != nil {
					t.Errorf("q.Pop: %v", err)
				}
				pops = append(pops, v)
			}

			for _, v := range pops {
				popCh <- v
			}
		}()
	}

	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			for j := i; j < pushCount; j += 20 {
				if err := q.Push(j + 1); err != nil {
					t.Errorf("q.Push: %v", err)
				}
			}
		}()
	}
	wg.Wait()
	close(popCh)

	// Remove the popped elements from the "want" set
	for v := range popCh {
		idx, found := slices.BinarySearch(wantElems, v)
		if !found {
			t.Errorf("%d not found", v)
		}
		wantElems = slices.Delete(wantElems, idx, idx+1)
	}

	var gotElems []int
	q.ForEach(func(elem int) {
		gotElems = append(gotElems, elem)
	})

	slices.Sort(gotElems)

	if len(gotElems) != pushCount-popCount {
		t.Errorf("expected %d elements, got %d", pushCount-popCount, len(gotElems))
	}

	if diff := cmp.Diff(wantElems, gotElems); diff != "" {
		t.Errorf("cmp.Diff: (-want, +got):\n%s", diff)
	}
}
