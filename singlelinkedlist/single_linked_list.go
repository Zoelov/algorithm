package singlelinkedlist

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	Head *Node[T]
	Size int
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Append(val T) {
	node := &Node[T]{
		Value: val,
	}
	if l.Head == nil {
		l.Head = node
		l.Size++
		return
	}

	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = node
	l.Size++
}

func (l *List[T]) HeadInsert(val T) {
	node := &Node[T]{
		Value: val,
		Next:  l.Head,
	}

	l.Head = node
	l.Size++
}
