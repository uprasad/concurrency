package queue

import "sync"

type BlockingQueue[T any] struct {
	buf  []T
	size int

	start, end int
	count      int

	mu    *sync.Mutex
	full  *sync.Cond
	empty *sync.Cond
}

func (q *BlockingQueue[T]) Push(elem T) error {
	q.mu.Lock()
	for q.count == q.size {
		q.full.Wait()
	}

	q.buf[q.end] = elem
	q.end = (q.end + 1) % q.size
	q.count++

	q.empty.Signal()
	q.mu.Unlock()
	return nil
}

func (q *BlockingQueue[T]) Pop() (T, error) {
	var v T
	q.mu.Lock()
	for q.count == 0 {
		q.empty.Wait()
	}

	v = q.buf[q.start]
	q.start = (q.start + 1) % q.size
	q.count--

	q.full.Signal()
	q.mu.Unlock()

	return v, nil
}

func (q *BlockingQueue[T]) Clear() {
	panic("not supported")
}

func (q *BlockingQueue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.count
}

func (q *BlockingQueue[T]) ForEach(f func(elem T)) {
	curr := q.start
	for i := 0; i < q.count; i++ {
		f(q.buf[curr])
		curr = (curr + 1) % q.size
	}
}

func NewBlockingQueue[T any](size int) *BlockingQueue[T] {
	var mu sync.Mutex
	return &BlockingQueue[T]{
		buf:  make([]T, size),
		size: size,

		mu:    &mu,
		full:  sync.NewCond(&mu),
		empty: sync.NewCond(&mu),
	}
}
