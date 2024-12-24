package queue

import "errors"

var (
	ErrEmpty = errors.New("queue is empty")
	ErrFull  = errors.New("queue is full")
)

type Queue[T any] interface {
	Push(elem T) error
	Pop() (T, error)
	Clear()
	Len() int
	ForEach(func(elem T))
}
