package main

import (
	"fmt"
	store "pl-project/storeLibs"
	"strconv"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup //WaitGroup that prevents main() from prematurely terminating goroutines
	minutes := 30         //total runtime
	storeOpen := true
	var waitLine = store.MakeQueue()    //Line customers wait in for a register to be open
	var register = []*(store.Register){ //array of registers
		store.MakeRegister(0, 0, 1, false),
		store.MakeRegister(1, 0, 1, false),
		store.MakeRegister(2, 0, 1, false),
	}
	var channel = []chan *(store.Customer){ //array of channels to send customers to registers
		make(chan *(store.Customer)),
		make(chan *(store.Customer)),
		make(chan *(store.Customer)),
	}
	wg.Add(4) //Adds active goroutines to the WaitGroup to prevent main() from terminating them
	/*
		Wait Line
		Register 0
		Register 1
		Register 2
	*/

	//goroutine to generate customers for the waiting line
	go func() {
		for minute := 0; minute < 60*minutes; minute++ {
			if store.IsCustomerAdded() {
				customer := store.MakeCustomer()
				fmt.Println("New Customer is in the wait line!", customer.ToString())
				waitLine.Enqueue(customer)
				//Checks to see if a register is open, and sends the first customer via a channel
				select {

				case channel[0] <- waitLine.Peek():
					fmt.Println("Sent customer " + strconv.Itoa(waitLine.Peek().ID) + " to register 0.")
					waitLine.Dequeue()
				case channel[1] <- waitLine.Peek():
					fmt.Println("Sent customer " + strconv.Itoa(waitLine.Peek().ID) + " to register 1.")
					waitLine.Dequeue()
				case channel[2] <- waitLine.Peek():
					fmt.Println("Sent customer " + strconv.Itoa(waitLine.Peek().ID) + " to register 2.")
					waitLine.Dequeue()
				default:
					fmt.Println("No open register")
				}
			}
		}
		//Finishes processing leftover customers after the time is up
		for !waitLine.IsEmpty() {
			select {
			case channel[0] <- waitLine.Peek():
				fmt.Println("Sent customer " + strconv.Itoa(waitLine.Peek().ID) + "to register 0.")
				waitLine.Dequeue()
			case channel[1] <- waitLine.Peek():
				fmt.Println("Sent customer " + strconv.Itoa(waitLine.Peek().ID) + "to register 1.")
				waitLine.Dequeue()
			case channel[2] <- waitLine.Peek():
				fmt.Println("Sent customer " + strconv.Itoa(waitLine.Peek().ID) + " to register 2.")
				waitLine.Dequeue()
			default:
				fmt.Println("No open register")
			}
		}
		storeOpen = false
		wg.Done()
	}()

	//Anonymous function for register[0]
	go func() {
		for !waitLine.IsEmpty() || storeOpen {
			select {
			case nextCustomer := <-channel[0]:
				register[0].Line.Enqueue(nextCustomer)
			case <-time.After(1 * time.Second):
				//register times out
			}
			for register[0].Line.Peek() != nil && !register[0].Line.Peek().HasBeenServiced() {
				if register[0].Line.Peek().Service() {
					fmt.Println("Customer with ID", register[0].Line.Peek().ID, "has been serviced at register 0.")
					register[0].Money.Add(store.Price(register[0].Line.Peek().Items))
					register[0].Line.Dequeue()
					register[0].TotalCustomersServiced++
				}
			}
		}
		wg.Done() //signals register is done servicing
	}()

	//Anonymous function for register[1]
	go func() {
		for !waitLine.IsEmpty() || storeOpen {
			select {
			case nextCustomer := <-channel[1]:
				register[1].Line.Enqueue(nextCustomer)
			case <-time.After(1 * time.Second):
				//register times out
			}
			for register[1].Line.Peek() != nil && !register[1].Line.Peek().HasBeenServiced() {
				if register[1].Line.Peek().Service() {
					fmt.Println("Customer with ID", register[1].Line.Peek().ID, "has been serviced at register 1.")
					register[1].Money.Add(store.Price(register[1].Line.Peek().Items))
					register[1].Line.Dequeue()
					register[1].TotalCustomersServiced++
				}
			}
		}
		wg.Done() //signals register is done servicing
	}()

	//Anonymous function for register[2]
	go func() {
		for !waitLine.IsEmpty() || storeOpen {
			select {
			case nextCustomer := <-channel[2]:
				register[2].Line.Enqueue(nextCustomer)
			case <-time.After(1 * time.Second):
				//register times out
			}
			for register[2].Line.Peek() != nil && !register[2].Line.Peek().HasBeenServiced() {
				if register[2].Line.Peek().Service() {
					fmt.Println("Customer with ID", register[2].Line.Peek().ID, "has been serviced at register 2.")
					register[2].Money.Add(store.Price(register[2].Line.Peek().Items))
					register[2].Line.Dequeue()
					register[2].TotalCustomersServiced++
				}
			}
		}
		wg.Done() //signals register is done servicing
	}()

	wg.Wait() //waits until all goroutines are finished before continuing

	fmt.Println("\n\nFinished running for", minutes, "minutes!")
	fmt.Println("In this time, register", register[0].ID, "serviced", register[0].TotalCustomersServiced, "customers, and made", register[0].Money.ToString())
	fmt.Println("In this time, register", register[1].ID, "serviced", register[1].TotalCustomersServiced, "customers, and made", register[1].Money.ToString())
	fmt.Println("In this time, register", register[2].ID, "serviced", register[2].TotalCustomersServiced, "customers, and made", register[2].Money.ToString())
}
