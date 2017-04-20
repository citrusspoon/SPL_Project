package storeLibs

import "fmt"

type Queue struct {
	First, Last *Customer
	Size        int
}

func MakeQueue(s int) *Queue {
	return &Queue{nil, nil, s}
}

func (queue *Queue) IsEmpty() bool {
	return queue.First == nil
}

func (queue *Queue) Enqueue(c *Customer) {
	if queue.IsEmpty() {
		queue.First = c
		queue.Last = c
	} else {
		queue.Last.NextCustomer = c
		queue.Last = c
	}

	queue.Size++
}

func (queue *Queue) Dequeue() {
	if queue.IsEmpty() {
		return
	}

	if queue.First == queue.Last {
		queue.Last = nil
	}

	queue.First = queue.First.NextCustomer

	queue.Size--
}

func (queue *Queue) Peek() *Customer {
	return queue.First
}

func (queue *Queue) PrintContents() {
	if queue.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	for customer := queue.First; customer != nil; customer = customer.NextCustomer {
		fmt.Println(*customer)
	}
}
