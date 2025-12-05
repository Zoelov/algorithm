package linked_list

import "fmt"

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}


func NewLinkedList[T any](head *Node[T]) *LinkedList[T] {
	return &LinkedList[T]{head: head}
}

func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}

func (l *LinkedList[T]) Build(data []T) {
	curr := l.head
	for i := range data {
		if curr == nil {
			if i == 0 {
				curr = &Node[T]{Data: data[0]}
				l.head = curr
			}
		} else {
			curr.Next = &Node[T]{Data: data[i]}
			curr = curr.Next
		}
	}
}

func (l *LinkedList[T]) Append(data T) {
	curr := l.head
	var pre *Node[T]
	for curr != nil {
		pre = curr
		curr = curr.Next
	}

	if pre == nil {
		pre = &Node[T]{Data: data}
		l.head = pre
	} else {
		pre.Next = &Node[T]{Data: data}
	}
}

func (l *LinkedList[T]) HeadInsert(data []T) {
	for i := range data {
		if l.head == nil {
			l.head = &Node[T]{Data: data[i]}
		} else {
			curr := &Node[T]{Data: data[i]}
			curr.Next = l.head
			l.head = curr
		}
	}
}

func (l *LinkedList[T]) Print() {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.Data)
		fmt.Print(" ")
		curr = curr.Next
	}
	fmt.Println()
}

func (l *LinkedList[T]) Get() []T {
	ret := make([]T, 0)
	curr := l.head
	for curr != nil {
		ret = append(ret, curr.Data)
		curr = curr.Next
	}

	return ret
}
