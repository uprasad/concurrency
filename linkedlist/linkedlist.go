package linkedlist

type node[T any] struct {
	elem T
	next *node[T]
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
