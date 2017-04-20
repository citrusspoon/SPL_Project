package main

// https://golang.org/doc/code.html
// http://www.newthinktank.com/2015/02/go-programming-tutorial/

import (
	"fmt"
	store "pl-project/storeLibs"
	"sync"
	//"reflect"
)


/*
Stuff to do
----------------
- Enforce max line size
- Create queue of people waiting to move to a non-full line
- Try to implement channels and a select statement to facilitate the above task
- Possibly move customer generation out of the register goroutine and into a "waiting line" goroutine
- Use channels to signal when time is over to stop generating customers?


*/

func main() {

	var wg sync.WaitGroup
	minutes := 10
	var register0 = store.MakeRegister(0, 0, 6, false)
	var register1 = store.MakeRegister(1, 0, 6, false)
	//var register2 = store.MakeRegister(2, 0, 6, false)

	
	wg.Add(2) //Adds active registers to the WaitGroup to prevent main() from terminating them

	//Anonymous function for register0
	go func() {
        
		
		for minute := 0; minute < 60*minutes; minute++ {
			if store.IsCustomerAdded() {
				customer := store.MakeCustomer()
				fmt.Println("New Customer is in the line!", customer.ToString())
				register0.Line.Enqueue(customer)
			}

			if register0.Line.Peek() != nil && register0.Line.Peek().Service() {
				fmt.Println("Customer with ID", register0.Line.Peek().ID, "has been serviced.")
				register0.Money.Add(store.Price(register0.Line.Peek().Items))
				register0.Line.Dequeue()
			}
		}
		wg.Done() //signals register is done servicing
    }()

	go func() {
        
		
		for minute := 0; minute < 60*minutes; minute++ {
			if store.IsCustomerAdded() {
				customer := store.MakeCustomer()
				fmt.Println("New Customer is in the line!", customer.ToString())
				register1.Line.Enqueue(customer)
			}

			if register1.Line.Peek() != nil && register1.Line.Peek().Service() {
				fmt.Println("Customer with ID", register1.Line.Peek().ID, "has been serviced.")
				register1.Money.Add(store.Price(register1.Line.Peek().Items))
				register1.Line.Dequeue()
			}
		}
		wg.Done() //signals register is done servicing
    }()


	wg.Wait() //waits until all goroutines are finished before continuing

	fmt.Println("\n\nFinished running for", minutes, "minutes!")
	fmt.Println("In this time, register", register0.ID, "made", register0.Money.ToString())
	fmt.Println("In this time, register", register1.ID, "made", register1.Money.ToString())

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