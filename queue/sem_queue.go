package queue

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

type SemQueue[T any] struct {
	buf  []T
	size int

	start, end int
	count      int

	mu    sync.Mutex
	full  *semaphore.Weighted
	empty *semaphore.Weighted
}

func (q *SemQueue[T]) Push(elem T) error {
	if err := q.full.Acquire(context.TODO(), 1); err != nil {
		panic(err)
	}
	q.mu.Lock()

	q.buf[q.end] = elem
	q.end = (q.end + 1) % q.size
	q.count++

	q.mu.Unlock()
	q.empty.Release(1)
	return nil
}

func (q *SemQueue[T]) Pop() (T, error) {
	var elem T

	if err := q.empty.Acquire(context.TODO(), 1); err != nil {
		panic(err)
	}
	q.mu.Lock()

	elem = q.buf[q.start]
	q.start = (q.start + 1) % q.size
	q.count--

	q.mu.Unlock()
	q.full.Release(1)

	return elem, nil
}

func (q *SemQueue[T]) Clear() {
	panic("not supported")
}

func (q *SemQueue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.count
}

func (q *SemQueue[T]) ForEach(f func(elem T)) {
	curr := q.start
	for i := 0; i < q.count; i++ {
		f(q.buf[curr])
		curr = (curr + 1) % q.size
	}
}

func NewSemQueue[T any](size int) *SemQueue[T] {
	empty := semaphore.NewWeighted(int64(size))
	if err := empty.Acquire(context.TODO(), int64(size)); err != nil {
		panic(err)
	}
	return &SemQueue[T]{
		buf:  make([]T, size),
		size: size,

		full:  semaphore.NewWeighted(int64(size)),
		empty: empty,
	}
}
