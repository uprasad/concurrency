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

type BasicLinkedList[T any] struct {
	head *node[T]
}

func (ll *BasicLinkedList[T]) Insert(elem T) {
	if ll.head == nil {
		ll.head = newNode(elem)
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = newNode(elem)
}

func (ll *BasicLinkedList[T]) ForEach(f func(elem T)) {
	for cur := ll.head; cur != nil; cur = cur.next {
		f(cur.elem)
	}
}

func (ll *BasicLinkedList[T]) Reset() { ll.head = nil }

func NewBasicLinkedList[T any]() *BasicLinkedList[T] {
	return &BasicLinkedList[T]{head: nil}
}
