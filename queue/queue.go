package queue

import "errors"

var ErrEmpty = errors.New("queue is empty")

type Queue[T any] interface {
	Push(elem T)
	Pop() (T, error)
	Clear()
	Len() int
	ForEach(func(elem T))
}
