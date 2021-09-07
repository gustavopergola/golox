package list

import (
	"fmt"
)

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
