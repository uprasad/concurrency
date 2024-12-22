package queue

type BasicQueue[T any] struct {
	buf []T
}

func (q *BasicQueue[T]) Push(elem T) {
	q.buf = append(q.buf, elem)
}

func (q *BasicQueue[T]) Pop() (T, error) {
	var v T
	if len(q.buf) == 0 {
		return v, ErrEmpty
	}

	v = q.buf[0]
	q.buf = q.buf[1:]

	return v, nil
}

func (q *BasicQueue[T]) Clear() {
	q.buf = nil
}

func (q *BasicQueue[T]) ForEach(f func(elem T)) {
	for _, elem := range q.buf {
		f(elem)
	}
}

func (q *BasicQueue[T]) Len() int {
	return len(q.buf)
}

func NewBasicQueue[T any]() *BasicQueue[T] {
	return &BasicQueue[T]{
		buf: make([]T, 0),
	}
}
