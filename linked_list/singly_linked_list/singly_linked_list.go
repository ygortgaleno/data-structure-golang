package singlylinkedlist

import "fmt"

type Node struct {
	Id   uint
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
}

func CreateList(ids ...uint) *SinglyLinkedList {
	l := &SinglyLinkedList{}
	n := &Node{}
	for idx, id := range ids {
		if idx == 0 {
			l.Head = n
		}

		n.Id = id
		if idx < len(ids)-1 {
			n.Next = &Node{}
			n = n.Next
		}
	}

	l.Tail = n

	return l
}

func (l *SinglyLinkedList) Print() {
	iterator := l.Head
	for iterator != nil {
		fmt.Println(*iterator)
		iterator = iterator.Next
	}

	fmt.Println("----------------")
}

func printReverse(currentNode *Node) {
	if currentNode == nil {
		return
	}

	printReverse(currentNode.Next)
	fmt.Println(*currentNode)
}

func (l *SinglyLinkedList) PrintReverse() {
	printReverse(l.Head)
	fmt.Println("----------------")
}

func (l *SinglyLinkedList) SearchElement(id uint) *Node {
	iterator := l.Head
	for iterator != nil {
		if iterator.Id == id {
			return iterator
		}

		iterator = iterator.Next
	}

	return nil
}

func (l *SinglyLinkedList) InsertEnd(id uint) {
	l.Tail.Next = &Node{
		Id: id,
	}

	l.Tail = l.Tail.Next
}

func (l *SinglyLinkedList) InsertBefore(id uint, n *Node) {
	if l.Head.Id == id {
		n.Next = l.Head
		l.Head = n
		return
	}

	prev := l.Head
	iterator := l.Head.Next
	for iterator != l.Tail {
		if iterator.Id == id {
			prev.Next = n
			n.Next = iterator
			return
		}

		prev = iterator
		iterator = iterator.Next
	}
}

func (l *SinglyLinkedList) RemoveElement(id uint) {
	if l.Head.Id == id {
		l.Head = l.Head.Next
		return
	}

	iterator := l.Head.Next
	prev := l.Head

	for iterator != l.Tail {
		if iterator.Id == id {
			prev.Next = iterator.Next

			if iterator.Next == l.Tail {
				l.Tail = iterator.Next
			}

			return
		}

		prev = iterator
		iterator = iterator.Next
	}
}
