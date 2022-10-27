package linkedlist

import (
	"fmt"
	"testing"
)

func Test_LinkedList(t *testing.T) {
	l := linkedlist[string]{}

	l.Add("faiz")
	l.Add("mohammad")
	l.Add("hafidza")

	fmt.Println(l)
	fmt.Println(l.Get("hafidza"))

	l.Remove("faiz")
	fmt.Println(l)
}
