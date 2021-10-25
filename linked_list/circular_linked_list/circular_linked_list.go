package circularLinkedList

import "fmt"

type Node struct {
	Id   uint
	Prev *Node
	Next *Node
}

type CircularLinkedList struct {
	Head *Node
	Tail *Node
}

func CreateList(ids ...uint) *CircularLinkedList {
	l := &CircularLinkedList{}
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

	l.Head.Prev = l.Tail
	l.Tail.Next = l.Head

	return l
}

func (l *CircularLinkedList) Print() {
	iterator := l.Head

	for iterator != l.Tail {
		fmt.Println(*iterator)
		iterator = iterator.Next
	}

	fmt.Println(*iterator)
}

func (l *CircularLinkedList) PrintReverse() {
	iterator := l.Tail

	for iterator != l.Head {
		fmt.Println(*iterator)
		iterator = iterator.Prev
	}

	fmt.Println(*iterator)
}

func (l *CircularLinkedList) SearchElement(id uint) *Node {
	if l.Tail.Id == id {
		return l.Tail
	}

	iterator := l.Head
	for iterator != l.Tail {
		if iterator.Id == id {
			return iterator
		}

		iterator = iterator.Next
	}

	return nil
}

func (l *CircularLinkedList) InsertEnd(id uint) {
	n := &Node{
		Id:   id,
		Prev: l.Tail,
		Next: l.Head,
	}

	l.Tail.Next = n
	l.Head.Prev = n
}

func (l *CircularLinkedList) InsertBefore(id uint, n *Node) {
	if l.Head.Id == id {
		n.Prev = l.Tail
		n.Next = l.Head.Next
		l.Head.Prev = n
		l.Tail.Next = n
		l.Head = n
		return
	}

	if l.Tail.Id == id {
		n.Prev = l.Tail.Prev
		n.Next = l.Tail
		l.Tail.Prev.Next = n
		l.Tail.Prev = n
		l.Tail = n
		return
	}

	iterator := l.Head
	for iterator != l.Tail {
		if iterator.Id == id {
			n.Prev = iterator.Prev
			n.Next = iterator
			iterator.Prev.Next = n
			iterator.Prev = n
			return
		}

		iterator = iterator.Next
	}
}

func (l *CircularLinkedList) RemoveElement(id uint) {
	if l.Head.Id == id {
		l.Head.Next.Prev = l.Head.Prev
		l.Head.Prev.Next = l.Head.Next
		l.Head = l.Head.Next
		return
	}

	if l.Tail.Id == id {
		l.Tail.Prev.Next = l.Tail.Next
		l.Tail.Next.Prev = l.Tail.Prev
		l.Tail = l.Tail.Prev
		return
	}

	iterator := l.Head
	for iterator != l.Tail {
		if iterator.Id == id {
			iterator.Next.Prev = iterator.Prev
			iterator.Prev.Next = iterator.Next
			return
		}
	}
}
