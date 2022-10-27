package linkedlist

import (
	"fmt"
	"sync"
)

func NewLinkedList[T comparable]() LinkedList[T] {
	return &linkedlist[T]{}
}

func (l *linkedlist[T]) Add(value T) {
	mu := sync.Mutex{}

	mu.Lock()
	newNode := new(node[T])
	newNode.value = value

	if l.head != nil {
		iteration := l.head
		for iteration.next != nil {
			iteration = iteration.next
		}

		iteration.next = newNode
	} else {
		l.head = newNode
	}
	mu.Unlock()
}

func (l linkedlist[T]) Get(value T) *node[T] {
	current := l.head
	for current != nil {
		if current.value == value {
			return current
		}

		current = current.next
	}

	return nil
}

func (l *linkedlist[T]) Remove(value T) {
	var (
		mu       sync.Mutex
		previous *node[T]
		current  = l.head
	)

	mu.Lock()
	for current != nil {
		if current == l.head && current.value == value {
			l.head = current.next
			return
		}

		if current.value == value {
			previous.next = current.next
			return
		}

		previous = current
		current = current.next
	}
	mu.Unlock()
}

func (l linkedlist[T]) String() string {
	var str string

	str = doString(l.head)

	return str
}

func doString[T comparable](n *node[T]) string {
	str := "{"

	if n != nil {
		str += " value: " + fmt.Sprint(n.value) + ", next: "
		if n.next != nil {
			str += doString(n.next)
		} else {
			str += fmt.Sprint(n.next)
		}

		str += " "
	}

	str += "}"

	return str
}
