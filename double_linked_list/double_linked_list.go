package double_linked_list

import "errors"

type DNode[T any] struct {
	Val  T
	Pre  *DNode[T]
	Next *DNode[T]
}

type DoubleLinkedList[T any] struct {
	dummpyHead *DNode[T] // head sentinel
	dummpyTail *DNode[T] // tail sentinel
	size       int
}

func NewDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	head := &DNode[T]{}
	tail := &DNode[T]{}

	head.Next = tail
	tail.Pre = head

	return &DoubleLinkedList[T]{
		dummpyHead: head,
		dummpyTail: tail,
	}
}

func (d *DoubleLinkedList[T]) Prepend(val T) {
	first := d.dummpyHead.Next
	node := &DNode[T]{
		Val:  val,
		Next: first,
		Pre:  d.dummpyHead,
	}

	d.dummpyHead.Next = node
	first.Pre = node
	d.size++
}

func (d *DoubleLinkedList[T]) Append(val T) {
	tail := d.dummpyTail.Pre
	node := &DNode[T]{
		Val:  val,
		Next: d.dummpyTail,
		Pre:  tail,
	}

	tail.Next = node
	d.dummpyTail.Pre = node
	d.size++
}

func (d *DoubleLinkedList[T]) Size() int {
	return d.size
}

func (d *DoubleLinkedList[T]) InsertAfter(node *DNode[T], val T) error {
	if node == nil || node == d.dummpyTail {
		return errors.New("node is invalide")
	}
	next := node.Next

	newNode := &DNode[T]{
		Val:  val,
		Next: next,
		Pre:  node,
	}
	node.Next = newNode
	next.Pre = newNode

	d.size++
	return nil
}

func (d *DoubleLinkedList[T]) InsertBefore(node *DNode[T], val T) error {
	if node == nil || node == d.dummpyHead {
		return errors.New("node is invalide")
	}

	pre := node.Pre
	newNode := &DNode[T]{
		Val:  val,
		Next: node,
		Pre:  pre,
	}
	pre.Next = newNode
	node.Pre = newNode
	d.size++

	return nil
}

func (d *DoubleLinkedList[T]) RemoveHead() (T, error) {
	var zero T
	if d.size == 0 {
		return zero, errors.New("list is empty")
	}

	head := d.dummpyHead.Next
	d.dummpyHead.Next = head.Next
	head.Next.Pre = d.dummpyHead

	head.Pre = nil
	head.Next = nil
	d.size--

	return head.Val, nil
}

func (d *DoubleLinkedList[T]) RemoveTail() (T, error) {
	var zero T
	if d.size == 0 {
		return zero, errors.New("list is empty")
	}

	tail := d.dummpyTail.Pre
	d.dummpyTail.Pre = tail.Pre
	tail.Pre.Next = d.dummpyTail

	tail.Next = nil
	tail.Pre = nil
	d.size--

	return tail.Val, nil
}

func (d *DoubleLinkedList[T]) Remove(node *DNode[T]) error {
	if node == nil || node == d.dummpyHead || node == d.dummpyTail {
		return errors.New("node is invalide")
	}

	pre := node.Pre
	next := node.Next

	pre.Next = next
	next.Pre = pre

	node.Pre = nil
	node.Next = nil
	d.size--

	return nil
}

func (d *DoubleLinkedList[T]) MoveToTail(node *DNode[T]) error {
	if node == nil || node == d.dummpyHead || node == d.dummpyTail {
		return errors.New("node is invalide")
	}

	if err := d.Remove(node); err != nil {
		return err
	}

	tail := d.dummpyTail.Pre
	tail.Next = node
	node.Next = d.dummpyTail
	node.Pre = tail
	d.dummpyTail.Pre = node
	d.size++

	return nil
}
