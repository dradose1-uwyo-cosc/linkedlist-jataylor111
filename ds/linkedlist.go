package ds

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

var emptyNode = &Node{"", nil}

// insert at the tail
func (l *LinkedList) Insert(value string) {
	node := &Node{value, nil}
	if l.Head == nil {
		l.Head = node
	} else {
		n := l.Head
		for n.next != nil {
			n = n.next
		}
		n.next = node
	}
	l.Size += 1
}

// inserts at a position, returns true if success or false if not, like if pos doesn't exist
func (l *LinkedList) InsertAt(position int, value string) error {
	new := &Node{value, nil}
	if position == 0 {
		new.next = l.Head
		l.Head = new
		l.Size++
		return nil
	}
	n := l.Head
	for c := 0; c < position-1; c++ {
		n = n.next
	}
	if n == nil {
		return fmt.Errorf("Error node not found")
	}
	new.next = n.next
	n.next = new
	l.Size++
	return nil
}

// remove first occurrence of the value
func (l *LinkedList) Remove(value string) error {
	n := l.Head
	if n.data == value {
		temp := n.next
		l.Head = temp
		l.Size--
		return nil
	} else {
		for n.next != nil {
			if n.next.data == value {
				n.next = n.next.next
				l.Size--
				return nil
			} else {
				n = n.next
			}
		}
		return fmt.Errorf("%s not found in list", value)
	}
}

// remove all occurrences of a value
func (l *LinkedList) RemoveAll(value string) {
	for l.Head.data == value {
		l.Head = l.Head.next
		l.Size--
	}
	n := l.Head
	for n != nil && n.next != nil {
		if n.next.data == value {
			n.next = n.next.next
			l.Size--
		} else {
			n = n.next
		}
	}
}

// remove at a position, if index exists
func (l *LinkedList) RemoveAt(pos int) error {
	if pos == 0 {
		l.Head = l.Head.next
		l.Size--
		return nil
	}
	n := l.Head
	for i := 0; i < pos-1; i++ {
		n = n.next
	}
	n.next = n.next.next
	l.Size--
	return nil
}

// checks if the linked list is empty
func (l *LinkedList) IsEmpty() bool {
	if l.Size != 0 {
		return false
	}
	return true
}

// get size of ll
func (l *LinkedList) GetSize() int {
	return l.Size
}

// reverse the list
func (l *LinkedList) Reverse() {
	n := l.Head
	var p *Node = nil
	for n != nil {
		next := n.next
		n.next = p
		p = n
		n = next
	}
	l.Head = p
}

// print the list
func (l *LinkedList) PrintList() {
	n := l.Head
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}
