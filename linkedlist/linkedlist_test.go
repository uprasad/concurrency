package linkedlist_test

import (
	"concurrency/linkedlist"
	"testing"

	"github.com/google/go-cmp/cmp"
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
