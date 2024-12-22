package linkedlist

import "sync"

type node[T any] struct {
	elem T
	next *node[T]

	mu sync.Mutex // used by coupled locking implementation
}

func newNode[T any](elem T) *node[T] {
	return &node[T]{
		elem: elem,
		next: nil,
	}
}

type LinkedList[T any] interface {
	Insert(elem T)
	ForEach(func(elem T))
	Reset()
}
