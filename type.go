package linkedlist

type LinkedList[T comparable] interface {
	Add(value T)
	Get(value T) *node[T]
	Remove(value T)
	String() string
}

type linkedlist[T comparable] struct {
	len  int
	head *node[T]
}

type node[T comparable] struct {
	value T
	next  *node[T]
}
