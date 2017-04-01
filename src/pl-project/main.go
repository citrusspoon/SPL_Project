package main

// https://golang.org/doc/code.html
// http://www.newthinktank.com/2015/02/go-programming-tutorial/

import (
	"fmt"
	"math/rand"
	"time"
)

//import "customer.go"

// http://stackoverflow.com/questions/15049903/how-to-use-custom-packages-in-golang
// import "/pl-project/storeLibs" // Not working? according to above it should

type Register struct {
	id                 int
	queueSize          int
	maxQueueSize       int
	currentlyServicing bool
}

var currentId = 0

const MINSERVICEDURATION int = 60
const MAXSERVICEDURATION int = 600
const NUMOFREGISTERS int = 5

func main() {

	//var registers[NUMOFREGISTERS] register

	// showing how the queue works
	queue := MakeQueue()

	fmt.Println("Empty?: ", queue.isEmpty())

	for i := 0; i < 10; i++ {
		queue.enqueue(MakeCustomer()) // 10 Customers go in line with IDs 1-10
	}

	for j := 0; j < 4; j++ {
		queue.dequeue() // Customers 1-4 are serviced, line now has 5-10
	}

	queue.printContents()
	fmt.Println("Queue Size: ", queue.size)
}

type Customer struct {
	id                 int
	serviceDurationSec int
	nextCustomer       *Customer
}

func (customer *Customer) decServiceTime() {
	customer.serviceDurationSec--
}

func (customer Customer) hasBeenServiced() bool {
	return customer.serviceDurationSec <= 0
}

func MakeCustomer() *Customer {

	/*
		This is necessary to reset the seed that the random function uses. Otherwise the numbers generated would be the same every time
		as Go doesn't reset the seed on its own. Just using the time isn't enough because it runs too fast, so I threw in an extra
		random number as well. It seems to be somewhat random.
	*/
	s1 := rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000000)))
	r1 := rand.New(s1)

	currentId++ //for some reason this can't be passed as an argument
	return &Customer{currentId, MINSERVICEDURATION + r1.Intn(MAXSERVICEDURATION-MINSERVICEDURATION), nil}
}

type Queue struct {
	first, last *Customer
	size        int
}

func MakeQueue() *Queue {
	return &Queue{nil, nil, 0}
}

func (queue *Queue) isEmpty() bool {
	return queue.first == nil
}

func (queue *Queue) enqueue(c *Customer) {
	if queue.isEmpty() {
		queue.first = c
		queue.last = c
	} else {
		queue.last.nextCustomer = c
		queue.last = c
	}

	queue.size++
}

func (queue *Queue) dequeue() {
	if queue.isEmpty() {
		return
	}

	if queue.first == queue.last {
		queue.last = nil
	}

	queue.first = queue.first.nextCustomer

	queue.size--
}

func (queue *Queue) printContents() {
	if queue.isEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	for customer := queue.first; customer != nil; customer = customer.nextCustomer {
		fmt.Println(*customer)
	}
}
