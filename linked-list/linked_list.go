package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

var (
	ErrEmptyList = errors.New("empty list")
)

func NewList(args ...interface{}) *List {
	list := &List{}
	for _, item := range args {
		list.PushBack(item)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	newNode := &Node{
		Value: v,
		next:  nil,
		prev:  nil,
	}
	if l.IsEmpty() {
		l.head = newNode
		l.tail = newNode
	} else {
		originalHead := l.head
		originalHead.prev = newNode
		newNode.next = originalHead
	}
	l.head = newNode
}

func (l *List) PushBack(v interface{}) {
	newNode := &Node{
		Value: v,
		next:  nil,
		prev:  nil,
	}
	if l.IsEmpty() {
		l.head = newNode
		l.tail = newNode
	} else {
		originalTail := l.tail
		originalTail.next = newNode
		newNode.prev = originalTail
	}
	l.tail = newNode
}

func (l *List) PopFront() (interface{}, error) {
	if l.IsEmpty() {
		return nil, ErrEmptyList
	}
	if l.head == l.tail {
		// last item in list
		item := l.head.Value
		l.head = nil
		l.tail = nil
		return item, nil
	}
	oldHead := l.First()
	newHead := oldHead.next
	newHead.prev = nil
	l.head = newHead
	return oldHead.Value, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.IsEmpty() {
		return nil, ErrEmptyList
	}
	if l.head == l.tail {
		// last item in list
		item := l.head.Value
		l.head = nil
		l.tail = nil
		return item, nil
	}
	oldTail := l.Last()
	newTail := oldTail.prev
	newTail.next = nil
	l.tail = newTail
	return oldTail.Value, nil
}

func (l *List) Reverse() {
	reversedList := NewList()
	item := l.First()
	for item != nil {
		reversedList.PushFront(item.Value)
		item = item.next
	}
	*l = *reversedList
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) IsEmpty() bool {
	return l.head == nil && l.tail == nil
}

func (l *List) Print() {
	node := l.First()
	for node != nil {
		fmt.Println(node.Value)
		node = node.next
	}
}
