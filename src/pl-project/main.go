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
	var ch0 = make(chan *Customer)
	var ch1 = make(chan *Customer)

	
	/*
	timeout := make(chan bool, 1)
	go func() {
    	time.Sleep(1 * time.Second)
    	timeout <- true
	}()
*/



	wg.Add(3) //Adds active registers to the WaitGroup to prevent main() from terminating them


	//goroutine to generate customers for the waiting line 
	go func(){

		for minute := 0; minute < 60*minutes; minute++ {
			if store.IsCustomerAdded() {
				customer := store.MakeCustomer()
				fmt.Println("New Customer is in the wait line!", customer.ToString())
				waitLine.Line.Enqueue(customer)

				//Checks to see if a register is open, and sends the first customer via a channel
				select {

					case ch0 <- waitLine.Line.Peek():
						fmt.Println("Sent " + strconv.Itoa(waitLine.Line.Peek()) + "to register 0.")
						waitLine.Line.Dequeue()
					case ch1 <- waitLine.Line.Peek():
						fmt.Println("Sent " + strconv.Itoa(waitLine.Line.Peek()) + "to register 1.")
						waitLine.Line.Dequeue()
					default: 
						fmt.Println("No open register")
				}



			}

		}
		//Finishes processing leftover customers after the time is up
		for !waitLine.Line.IsEmpty(){

			select {

					case ch0 <- waitLine.Line.Peek():
						fmt.Println("Sent " + strconv.Itoa(waitLine.Line.Peek()) + "to register 0.")
						waitLine.Line.Dequeue()
					case ch1 <- waitLine.Line.Peek():
						fmt.Println("Sent " + strconv.Itoa(waitLine.Line.Peek()) + "to register 1.")
						waitLine.Line.Dequeue()
					default: 
						fmt.Println("No open register")
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