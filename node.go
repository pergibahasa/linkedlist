package linkedlist

import "fmt"

func (n node[T]) String() string {
	str := "{ "

	str += "value: " + fmt.Sprint(n.value) + " next: " + fmt.Sprint(n.next)

	str += " }"

	return str
}
