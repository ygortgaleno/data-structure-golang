package stack

import "fmt"

type Node struct {
	Id   uint
	Prev *Node
	Next *Node
}

type Stack struct {
	Head *Node
	Tail *Node
}

func CreateStack(ids ...uint) *Stack {
	s := &Stack{}
	n := &Node{}

	for idx, id := range ids {
		if idx == 0 {
			s.Head = n
		}

		n.Id = id
		n.Prev = s.Tail
		s.Tail = n

		if idx < len(ids)-1 {
			n.Next = &Node{}
			n = n.Next
		}
	}

	return s
}

func (s *Stack) Print() {
	for iterator := s.Head; iterator != nil; iterator = iterator.Next {
		fmt.Println(*iterator)
	}
}

func (s *Stack) Push(id uint) {
	if s.Tail != nil {
		s.Tail.Next = &Node{
			Id:   id,
			Prev: s.Tail,
		}
		s.Tail = s.Tail.Next
	}
}

func (s *Stack) Pop() {
	if s.Tail != nil {
		s.Tail.Prev.Next = nil
		s.Tail = s.Tail.Prev
	}
}
