package linkedlist

import "sync"

type LockLinkedList[T any] struct {
	head *node[T]

	mu sync.Mutex
}

func (ll *LockLinkedList[T]) Insert(elem T) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	if ll.head == nil {
		ll.head = newNode(elem)
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = newNode(elem)
}

func (ll *LockLinkedList[T]) ForEach(f func(elem T)) {
	for cur := ll.head; cur != nil; cur = cur.next {
		f(cur.elem)
	}
}

func (ll *LockLinkedList[T]) Reset() {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	ll.head = nil
}

func NewLockLinkedList[T any]() *LockLinkedList[T] {
	return &LockLinkedList[T]{head: nil}
}
