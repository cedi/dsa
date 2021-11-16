package linkedlist

import "fmt"

// LinkedList is the Datatype that holdes the HEAD Node to the linked list.
// on the LinkedList object you then call all the fancy methods I've implemented
type LinkedList struct {
	head *Node
}

// Node is the Data Node which constructs my LinkedList
type Node struct {
	Data int
	Next *Node
}

// New creates a empty LinkedList with no items
func New() *LinkedList {
	return &LinkedList{
		head: nil,
	}
}

// NewFromArgs takes a input string and converts it to a linked list
func NewFromArgs(items ...int) *LinkedList {
	list := New()
	for _, itm := range items {
		list.InsertBack(itm)
	}

	return list
}

// Len returns the length of the LinkedList
func (l *LinkedList) Len() int {
	len := 0
	for iter := l.head; iter != nil; iter = iter.Next {
		len++
	}
	return len
}

// String returns the string representatino of the LinkedList
func (l *LinkedList) String() string {
	output := ""

	for iter := l.head; iter != nil; iter = iter.Next {
		output += fmt.Sprintf("[%d]->", iter.Data)
	}

	output = output[:len(output)-2]

	return output
}

// GetHead will return the Head node
func (l *LinkedList) GetHead() *Node {
	return l.head
}

// InsertFront will insert a new Element at the beginning
// Example: InsertFront("Node 3") will cause:
// 	[Head] -> Node 1 -> Node 2 -> [EOF]
//	to become
//	[HEAD] -> Node 3 -> Node 1 -> Node 2 -> [EOF]
// Returns the freshly inserted Element
// Runtime complexity is O(1)
func (l *LinkedList) InsertFront(data int) *Node {
	return l.InsertAfter(nil, data)
}

// InsertBack will insert a new Element at the end
// Example: InsertBack("Node 3") will cause:
// 	[Head] -> Node 1 -> Node 2 -> [EOF]
//	to become
//	[HEAD] -> Node 1 -> Node 2 -> Node 3 -> [EOF]
// Returns the freshly inserted Element
// Runtime complexity is O(n)
func (l *LinkedList) InsertBack(data int) *Node {

	if l.head == nil {
		return l.InsertFront(data)
	}

	return l.InsertAfter(l.GetLast(), data)
}

// InsertAfter will insert a new Element after the given element
// Example: InsertAfter("Node 1", "Node 3") will cause:
// 	[Head] -> Node 1 -> Node 2 -> [EOF]
//	to become
//	[HEAD] -> Node 1 -> Node 3 -> Node 2 -> [EOF]
// Returns the freshly inserted Element
// Runtime complexity is O(1)
func (l *LinkedList) InsertAfter(node *Node, data int) *Node {
	var newNextNode *Node

	if node == nil {
		newNextNode = l.head
	} else {
		newNextNode = node.Next
	}

	newNode := &Node{
		Data: data,
		Next: newNextNode,
	}

	if node == nil {
		l.head = newNode
	} else {
		node.Next = newNode
	}
	return newNode
}

// InsertBefore will insert a new Element before the given element
// Example: InsertBefore("Node 2", "Node 3") will cause:
// 	[Head] -> Node 1 -> Node 2 -> [EOF]
//	to become
//	[HEAD] -> Node 1 -> Node 3 -> Node 2 -> [EOF]
// Returns the freshly inserted Element
// Runtime complexity is O(n)
func (l *LinkedList) InsertBefore(node *Node, data int) *Node {
	if node == nil {
		return l.InsertFront(data)
	}

	newNode := &Node{
		Data: data,
		Next: node,
	}

	var beforeNode *Node

	for iter := l.head; iter != nil; iter = iter.Next {
		if iter.Next == node {
			beforeNode = iter
			break
		}
	}

	if beforeNode == nil {
		return l.InsertFront(data)
	}

	beforeNode.Next = newNode

	return newNode
}

func (l *LinkedList) GetLast() *Node {
	lastElement := l.head

	for lastElement.Next != nil {
		lastElement = lastElement.Next
	}

	return lastElement
}

// Delete will delete a node from the Linked List
// Example: Delete("Node 2") will cause:
//	[HEAD] -> Node 1 -> Node 2 -> Node 3 -> [EOF]
//	to become
// 	[Head] -> Node 1 -> Node 3 -> [EOF]
// Returns the data of the deleted Node
func (l *LinkedList) Delete(node *Node) int {
	if node == l.head {
		l.head = node.Next
		return node.Data
	}

	prevNode := l.head
	for prevNode != nil && prevNode.Next != node {
		prevNode = prevNode.Next
	}

	if prevNode != nil {
		prevNode.Next = node.Next
	}

	return node.Data
}
