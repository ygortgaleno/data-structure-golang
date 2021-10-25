package doublylinkedlist

import "fmt"

type Node struct {
	Id   uint
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func CreateList(ids ...uint) *DoublyLinkedList {
	l := &DoublyLinkedList{}
	n := &Node{}
	for idx, id := range ids {
		if idx == 0 {
			l.Head = n
		}

		n.Id = id
		n.Prev = l.Tail
		l.Tail = n

		if idx < len(ids)-1 {
			n.Next = &Node{}
			n = n.Next
		}
	}

	return l
}

func (l *DoublyLinkedList) Print() {
	iterator := l.Head
	for iterator != nil {
		fmt.Println(*iterator)
		iterator = iterator.Next
	}

	fmt.Println("----------------")
}

func (l *DoublyLinkedList) PrintReverse() {
	iterator := l.Tail

	for iterator != nil {
		fmt.Println(*iterator)
		iterator = iterator.Prev
	}

	fmt.Println("----------------")
}

func (l *DoublyLinkedList) SearchElement(id uint) *Node {
	iterator := l.Head

	for iterator != nil {
		if iterator.Id == id {
			return iterator
		}

		iterator = iterator.Next
	}

	return nil
}

func (l *DoublyLinkedList) InsertEnd(id uint) {
	n := &Node{
		Id:   id,
		Prev: l.Tail,
	}

	l.Tail.Next = n
	l.Tail = n
}

func (l *DoublyLinkedList) InsertBefore(id uint, n *Node) {
	if l.Head.Id == id {
		n.Next = l.Head
		l.Head.Prev = n
		l.Head = n
		return
	}

	if l.Tail.Id == id {
		n.Prev = l.Tail
		l.Tail.Next = n
		l.Tail = n
		return
	}

	iterator := l.Head

	for iterator != nil {
		if iterator.Id == id {
			n.Next = iterator
			n.Prev = iterator.Prev
			iterator.Prev.Next = n
			iterator.Prev = n
			return
		}

		iterator = iterator.Next
	}
}

func (l *DoublyLinkedList) RemoveElement(id uint) {
	if l.Head.Id == id {
		l.Head.Next.Prev = nil
		l.Head = l.Head.Next
		return
	}

	if l.Tail.Id == id {
		l.Tail.Prev.Next = nil
		l.Tail = l.Tail.Prev
		return
	}

	iterator := l.Head

	for iterator != nil {
		if iterator.Id == id {
			iterator.Prev.Next = iterator.Next
			iterator.Next.Prev = iterator.Prev
			return
		}

		iterator = iterator.Next
	}
}
