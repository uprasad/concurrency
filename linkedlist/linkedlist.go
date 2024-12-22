package linkedlist

type LinkedList[T any] interface {
	Insert(elem T)
	ForEach(func(elem T))
	Reset()
}
