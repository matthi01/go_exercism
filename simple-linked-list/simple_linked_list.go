package linkedlist

import (
	"errors"
)

type List struct {
	head *Element
	size int
}

type Element struct {
	value int
	next  *Element
}

func New(items []int) *List {
	newList := &List{}
	for _, n := range items {
		newList.Push(n)
	}
	return newList
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	newElement := &Element{
		value: element,
		next:  nil,
	}
	if l.head == nil {
		l.head = newElement
		l.size++
		return
	}
	newElement.next = l.head
	l.head = newElement
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}
	poppedElement := l.head
	l.head = poppedElement.next
	l.size--
	return poppedElement.value, nil
}

func (l *List) Array() []int {
	return l.Reverse().Values()
}

func (l *List) Reverse() *List {
	return New(l.Values())
}

func (l *List) Values() []int {
	values := []int{}
	current := l.head
	for current != nil {
		values = append(values, current.value)
		current = current.next
	}
	return values
}
