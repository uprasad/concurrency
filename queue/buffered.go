package queue

type BufferedQueue[T any] struct {
	buf  chan T
	size int
}

func (q *BufferedQueue[T]) Push(elem T) error {
	select {
	case q.buf <- elem:
		return nil
	default:
		return ErrFull
	}
}

func (q *BufferedQueue[T]) Pop() (T, error) {
	var v T
	select {
	case v = <-q.buf:
		return v, nil
	default:
		return v, ErrEmpty
	}
}

func (q *BufferedQueue[T]) Clear() {
	q.buf = make(chan T, q.size)
}

func (q *BufferedQueue[T]) ForEach(f func(elem T)) {
	panic("unsupported")
}

func NewBufferedQueue[T any](size int) *BufferedQueue[T] {
	return &BufferedQueue[T]{
		buf:  make(chan T, size),
		size: size,
	}
}
