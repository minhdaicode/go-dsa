package linkedlist

import "fmt"

type Singly[T comparable] struct {
	Head *Node[T]
	len  int
}

func NewSingly[T comparable]() *Singly[T] {
	return &Singly[T]{}
}

func (l *Singly[T]) InsertHead(data T) {
	nn := NewNode(data)
	nn.Next = l.Head
	l.Head = nn
	l.len++
}

func (l *Singly[T]) InsertTail(data T) {
	if l.isEmpty() {
		l.InsertHead(data)
		return
	}

	cur := l.Head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = NewNode(data)
	l.len++
}

func (l *Singly[T]) Insert(data T, position int) error {
	switch {
	case position < 0 || position > l.len:
		return ErrPosOutOfRange
	case position == 0:
		l.InsertHead(data)
		return nil
	case position == l.len:
		l.InsertTail(data)
		return nil
	default:
		cur := l.Head
		prev := cur
		for i := 0; cur != nil; i++ {
			if i == position {
				nn := NewNode(data)
				prev.Next = nn
				nn.Next = cur
				l.len++
				break
			}
			prev = cur
			cur = cur.Next
		}
		return nil
	}
}

func (l *Singly[T]) Transverse() {
	for cur := l.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " -> ")
	}
	fmt.Print("nil")
}

func (l *Singly[T]) Len() int {
	return l.len
}

func (l *Singly[T]) isEmpty() bool {
	return l.Head == nil || l.len == 0
}
