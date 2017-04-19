package main

//https://www.youtube.com/watch?v=zbFDjCHzN50

/*

Goroutine General Info
======================================================================================
	- Any function can be made a goroutine by putting "go" in front when you call it
		ex. (  go fmtPrintln("stuff")  )
	- goroutines share the same memory 
	- If the main() method ends, it will terminate all executing goroutines
	- 
======================================================================================


WaitGroups
======================================================================================
	- A WaitGroup is an entity of the "sync" package that allows you to control goroutines using an internal counter
	- Creating a waitgroup:
		 var wg sync.WaitGroup
	- WaitGroup Functions:
		Add(int) -- Adds an int to the WaitGroup's internal counter
		Done() -- Subtracts 1 from the WaitGroup's internal counter
		Wait() -- stops execution of a function until the WaitGroup's internal counter reaches 0
	- A freshly initilized WaitGroup's counter starts at 0
	- Basically every time you start a goroutine, you should use Add() to raise the counter
	- Every time a goroutine is done you should call Done()
	- At the very least you need a WaitGroup to stop main() from killing all the goroutines prematurely

======================================================================================


*/
import (
	"fmt"
	"sync"
)

func main() {


}
