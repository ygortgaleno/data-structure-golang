package queue

import "fmt"

type Node struct {
	Id   uint
	Next *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func CreateQueue(ids ...uint) *Queue {
	q := &Queue{}
	n := &Node{}

	for idx, id := range ids {
		if idx == 0 {
			q.Head = n
		}

		n.Id = id
		q.Tail = n

		if idx < len(ids)-1 {
			n.Next = &Node{}
			n = n.Next
		}
	}

	return q
}

func (q *Queue) Print() {
	for iterator := q.Head; iterator != nil; iterator = iterator.Next {
		fmt.Println(*iterator)
	}
}

func (q *Queue) Push(id uint) {
	q.Tail.Next = &Node{
		Id: id,
	}
	q.Tail = q.Tail.Next
}

func (q *Queue) Pop() {
	if q.Head != nil {
		q.Head = q.Head.Next
	}
}
