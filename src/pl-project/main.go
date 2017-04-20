package main

// https://golang.org/doc/code.html
// http://www.newthinktank.com/2015/02/go-programming-tutorial/

import (
	"fmt"
	store "pl-project/storeLibs"
	"sync"
	"time"
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
	var waitLine = store.MakeQueue();
	var register0 = store.MakeRegister(0, 0, 6, false)	
	var register1 = store.MakeRegister(1, 0, 6, false)
	//var register2 = store.MakeRegister(2, 0, 6, false)
	var reg0ch = make(chan *Customer)
	var reg1ch = make(chan *Customer)

	
	wg.Add(3) //Adds active registers to the WaitGroup to prevent main() from terminating them



	go func(){

		for minute := 0; minute < 60*minutes; minute++ {
			if store.IsCustomerAdded() {
				customer := store.MakeCustomer()
				fmt.Println("New Customer is in the line!", customer.ToString())
				waitLine.Line.Enqueue(customer)

				select {	
				case x := <-reg0ch
					reg0ch <- waitLine.Line.Peek()
					waitLine.Line.Dequeue()
				case x := <-reg1ch
					reg1ch <- waitLine.Line.Peek()
					waitLine.Line.Dequeue()
				default:
					fmt.Println("Waiting for open register")	
				}





			}

		}
		

		for !waitLine.IsEmpty() {

			select {	
				case x := <-reg0ch
					reg0ch <- waitLine.Line.Peek()
					waitLine.Line.Dequeue()
				case x := <-reg1ch
					reg1ch <- waitLine.Line.Peek()
					waitLine.Line.Dequeue()
				default:
					fmt.Println("Waiting for open register")	
				}



		}



		wg.Done()
	}()









	//Anonymous function for register0
	go func() {
        
		
		
		wg.Done() //signals register is done servicing
    }()

	go func() {
        
		
	
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









/*
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






*/