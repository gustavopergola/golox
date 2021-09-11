package list

import (
	"errors"
	"fmt"
)

type Node struct {
	Prev  *Node
	Next  *Node
	Value string
}

type List struct {
	Length int
	Start  *Node
	End    *Node
}

func (l *List) Insert(newValue string) {
	var newNode Node = Node{Value: newValue}

	l.Length++
	if l.Length == 1 {
		l.Start = &newNode
		l.End = &newNode
	} else {
		last := l.End
		last.Next = &newNode
		newNode.Prev = last
		l.End = &newNode
	}
}

func (l *List) Print() {
	node := l.Start
	for {
		if node == nil {
			break
		}

		fmt.Println(node.Value)
		node = node.Next
	}
	fmt.Printf("Total of %d items\n", l.Length)
}

func (l *List) Find(value string) *Node {
	node := l.Start
	for {
		if node == nil || node.Value == value {
			return node
		}

		node = node.Next
	}
}

func (l *List) Delete(value string) error {
	node := l.Find(value)

	if node == nil {
		return errors.New("Node not found!")
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Prev == nil {
		l.Start = node.Next
	}
	if node.Next == nil {
		l.End = node.Prev
	}
	l.Length--

	return nil
}
