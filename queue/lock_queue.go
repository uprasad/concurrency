package queue

import "sync"

type LockQueue[T any] struct {
	buf []T

	mu sync.Mutex
}

func (q *LockQueue[T]) Push(elem T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.buf = append(q.buf, elem)
}

func (q *LockQueue[T]) Pop() (T, error) {
	var v T
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.buf) == 0 {
		return v, ErrEmpty
	}

	v = q.buf[0]
	q.buf = q.buf[1:]

	return v, nil
}

func (q *LockQueue[T]) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.buf = nil
}

func (q *LockQueue[T]) ForEach(f func(elem T)) {
	for _, elem := range q.buf {
		f(elem)
	}
}

func (q *LockQueue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.buf)
}

func NewLockQueue[T any]() *LockQueue[T] {
	return &LockQueue[T]{
		buf: make([]T, 0),
	}
}
