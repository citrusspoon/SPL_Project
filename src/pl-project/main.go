package main

// https://golang.org/doc/code.html
// http://www.newthinktank.com/2015/02/go-programming-tutorial/

import (
	"fmt"
	store "pl-project/storeLibs"
	//"reflect"
)

func main() {


	//var registers = []*Register{store.MakeRegister(), store.MakeRegister()}
	/*
	var register0 = store.MakeRegister()
	var register1 = store.MakeRegister()
	var register2 = store.MakeRegister()
	var register3 = store.MakeRegister()
	var register4 = store.MakeRegister()
	*/
/*
	for j := 0; j < 5; j++ {
	
		registers[j] = store.MakeRegister()
	
	}
*/
	register := store.MakeRegister()
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
			register.Money.Add(store.Price(queue.Peek().Items))
			queue.Dequeue()
		}
	}

	fmt.Println("\n\nFinished running for", minutes, "minutes!")
	fmt.Println("In this time, register", register.ID, "made", register.Money.ToString())
	//fmt.Println(reflect.TypeOf(register))
}
