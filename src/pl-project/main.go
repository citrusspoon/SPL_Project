package main

// https://golang.org/doc/code.html
// http://www.newthinktank.com/2015/02/go-programming-tutorial/

import (
	"fmt"
	store "pl-project/storeLibs"
)

func main() {

	queue := store.MakeQueue()

	minutes := 10

	for minute := 0; minute < 60*minutes; minute++ {
		if store.IsCustomerAdded() {
			customer := store.MakeCustomer()
			fmt.Println("New Customer is in the line!", customer.ToString())
			queue.Enqueue(customer)
		}

		if queue.Peek() != nil && queue.Peek().Service() {
			fmt.Println("Customer with ID", queue.Peek().ID, "has been serviced.")
			queue.Dequeue()
		}
	}
}
