package linkedlist

type CoupleLockLinkedList[T any] struct {
	head *node[T]
}

func (ll *CoupleLockLinkedList[T]) Insert(elem T) {
	cur := ll.head
	cur.mu.Lock()

	for ; cur.next != nil; cur = cur.next {
		cur.next.mu.Lock()
		cur.mu.Unlock()
	}

	cur.next = newNode(elem)
	cur.mu.Unlock()
}

func (ll *CoupleLockLinkedList[T]) ForEach(f func(elem T)) {
	// Skip the dummy head node when traversing
	for cur := ll.head.next; cur != nil; cur = cur.next {
		f(cur.elem)
	}
}

func (ll *CoupleLockLinkedList[T]) Reset() {
	ll.head.mu.Lock()
	defer ll.head.mu.Unlock()
	ll.head = &node[T]{next: nil}
}

func NewCoupleLockLinkedList[T any]() *CoupleLockLinkedList[T] {
	// Dummy head node needed for initial node lock when inserting / resetting
	dummy := &node[T]{next: nil}

	return &CoupleLockLinkedList[T]{
		head: dummy,
	}
}
