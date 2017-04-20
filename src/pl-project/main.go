package main

// https://golang.org/doc/code.html
// http://www.newthinktank.com/2015/02/go-programming-tutorial/

import (
	"fmt"
	store "pl-project/storeLibs"
	"sync"
	//"reflect"
)

func main() {

	var wg sync.WaitGroup
	var register = store.MakeRegister(0, 0, 6, false)
	

	wg.Add(1)
	go func() {
        
		minutes := 10
		for minute := 0; minute < 60*minutes; minute++ {
			if store.IsCustomerAdded() {
				customer := store.MakeCustomer()
				fmt.Println("New Customer is in the line!", customer.ToString())
				register.Line.Enqueue(customer)
			}

			if register.Line.Peek() != nil && register.Line.Peek().Service() {
				fmt.Println("Customer with ID", register.Line.Peek().ID, "has been serviced.")
				register.Money.Add(store.Price(register.Line.Peek().Items))
				register.Line.Dequeue()
			}
		}

		fmt.Println("\n\nFinished running for", minutes, "minutes!")
		fmt.Println("In this time, register", register.ID, "made", register.Money.ToString())
		wg.Done()
    }()


	wg.Wait()


}
/*

//I can't figure out how to pass the Register into this, so if someone else figures it out we can use this in main instead of anonymous functions to make it look better
func startRegister(){

	minutes := 10
	for minute := 0; minute < 60*minutes; minute++ {
		if store.IsCustomerAdded() {
			customer := store.MakeCustomer()
			fmt.Println("New Customer is in the line!", customer.ToString())
			register.Line.Enqueue(customer)
		}

		if register.Line.Peek() != nil && register.Line.Peek().Service() {
			fmt.Println("Customer with ID", register.Line.Peek().ID, "has been serviced.")
			register.Money.Add(store.Price(register.Line.Peek().Items))
			register.Line.Dequeue()
		}
	}

	fmt.Println("\n\nFinished running for", minutes, "minutes!")
	fmt.Println("In this time, register", register.ID, "made", register.Money.ToString())

}*/