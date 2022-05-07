package linkedlist

import (
	"errors"
	"fmt"
)

// Define List and Node types here.
type List struct {
	head *Node
	tail *Node
}
type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

var (
	ErrEmptyList = errors.New("empty list")
)

func NewList(args ...interface{}) *List {
	newList := &List{head: nil, tail: nil}
	for _, v := range args {
		fmt.Println("push...", v)
		newList.PushBack(v)
	}
	fmt.Println("first:", newList.First().Value)
	fmt.Println("last:", newList.Last().Value)
	return newList
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	newNode := Node{
		Value: v,
		next:  l.First(),
		prev:  nil,
	}
	// if new list
	if l.head == nil && l.tail == nil {
		fmt.Println("found new list")
		l.head = &newNode
		l.tail = &newNode
		return
	}
	fmt.Println("adding to list")
	newNode.prev = l.head
	l.head = &newNode
}

func (l *List) PushBack(v interface{}) {
	newNode := Node{
		Value: v,
		next:  nil,
		prev:  l.Last(),
	}
	if l.head == nil && l.tail == nil {
		fmt.Println("found new list")
		l.head = &newNode
		l.tail = &newNode
		return
	}
	fmt.Println("adding to list")
	newNode.next = l.tail
	l.tail = &newNode
}

func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		return "", ErrEmptyList
	}
	poppedNode := l.head
	poppedNode.prev = nil
	return poppedNode.Value, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
		return "", ErrEmptyList
	}
	poppedNode := l.tail
	poppedNode.next = nil
	return poppedNode.Value, nil
}

func (l *List) Reverse() {
	panic("Please implement the Reverse function")
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
