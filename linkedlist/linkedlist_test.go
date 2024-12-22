package linkedlist_test

import (
	"concurrency/linkedlist"
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	parallelisms = []int{1, 2, 4, 8, 16}
)

func TestBasicLinkedList(t *testing.T) {
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
			ll := linkedlist.NewBasicLinkedList[int]()
			for _, elem := range test.elems {
				ll.Insert(elem)
			}

			gotElems := make([]int, 0, len(test.elems))
			ll.ForEach(func(elem int) {
				gotElems = append(gotElems, elem)
			})

			if diff := cmp.Diff(test.elems, gotElems); diff != "" {
				t.Errorf("cmp.Diff: (-want, +got):\n%s", diff)
			}

			ll.Reset()
			var gotResetElems []int
			ll.ForEach(func(elem int) {
				gotResetElems = append(gotResetElems, elem)
			})
			if diff := cmp.Diff([]int(nil), gotResetElems); diff != "" {
				t.Errorf("cmp.Diff reset: (-want, +got)\n%s", diff)
			}
		})
	}
}

func BenchmarkBasicLinkedList_1_000(b *testing.B) {
	for _, p := range parallelisms {
		b.Run(fmt.Sprintf("%d", p), func(b *testing.B) {
			ll := linkedlist.NewBasicLinkedList[int]()
			benchmarkLinkedList(b, ll, p, 1_000)
			count := 0
			ll.ForEach(func(elem int) { count++ })
			// b.Logf("expected count: %d, got %d\n", 1_000*p, count)
		})
	}
}

func BenchmarkBasicLinkedList_10_000(b *testing.B) {
	for _, p := range parallelisms {
		b.Run(fmt.Sprintf("%d", p), func(b *testing.B) {
			ll := linkedlist.NewBasicLinkedList[int]()
			benchmarkLinkedList(b, ll, p, 10_000)
			count := 0
			ll.ForEach(func(elem int) { count++ })
			b.Logf("expected count: %d, got %d\n", 10_000*p, count)
		})
	}
}

func benchmarkLinkedList(b *testing.B, ll linkedlist.LinkedList[int], parallelism, count int) {
	elems := make([]int, 0, count)
	for i := 0; i < count; i++ {
		elems = append(elems, rand.Int())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.Reset()
		// Kick off goroutines that increment in parallel
		var wg sync.WaitGroup
		wg.Add(parallelism)
		for j := 0; j < parallelism; j++ {
			go func() {
				defer wg.Done()
				insertElems(ll, elems)
			}()
		}
		wg.Wait()
	}
}

func insertElems(ll linkedlist.LinkedList[int], elems []int) {
	for _, elem := range elems {
		ll.Insert(elem)
	}
}
