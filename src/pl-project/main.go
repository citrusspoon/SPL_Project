package main

// https://golang.org/doc/code.html
// http://www.newthinktank.com/2015/02/go-programming-tutorial/

import (
	"fmt"
	"math/rand"
	"strconv"
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

	minutes := 10

	for j := 0; j < 60*minutes; j++ {
		if isCustomerAdded() {
			customer := MakeCustomer()
			fmt.Println("New Customer is in the line!", customer.toString())
			queue.enqueue(customer)
		}

		if queue.peek() != nil && queue.peek().service() {
			fmt.Println("Customer with ID", queue.peek().id, "has been serviced.")
			queue.dequeue()
		}
	}

	queue.printContents()
	fmt.Println("Queue Size: ", queue.size, " peek ", queue.peek().serviceDurationSec)
}

func isCustomerAdded() bool {
	s := rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000000)))
	r := rand.New(s)

	// 1% chance someone is added per second
	return r.Intn(100) == 0
}

type Customer struct {
	id                 int
	serviceDurationSec int
	nextCustomer       *Customer
}

func (customer *Customer) decServiceTime() {
	customer.serviceDurationSec--
}

// returns true if the customer has finished being serviced, false otherwise
func (customer *Customer) service() bool {
	customer.decServiceTime()
	return customer.hasBeenServiced()
}

func (customer *Customer) hasBeenServiced() bool {
	return customer.serviceDurationSec <= 0
}

func (customer *Customer) toString() string {
	return "ID: " + strconv.Itoa(customer.id) + ", Service Time (seconds): " + strconv.Itoa(customer.serviceDurationSec)
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

func (queue *Queue) peek() *Customer {
	return queue.first
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
