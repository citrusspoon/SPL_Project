package main

//https://www.youtube.com/watch?v=zbFDjCHzN50
//https://golang.org/pkg/sync/#example_WaitGroup
//https://www.random.org/strings/?num=10&len=10&digits=on&unique=on&format=plain&rnd=new
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


Channels
======================================================================================
	- Conceptually these are typed pipes that connect goroutines for data sharing
	- By default they are unbuffered/synchronous
	- This means that when attempting to send/recieve over a channel, both sides must be ready for it to occur
	- If a goroutine tries to send data before another is ready to recieve it and vice versa, it will pause executing
	- Creating a channel:
		ch0 := make(chan int)
	- Sending data over a channel
		ch0 <- 5
	- Recieving data via a channel
		var x int
		x <-ch0 //pull an int and store in x
		<-ch0 //pull an int and discard it
		var y := <-ch0 //it can also be used like this
	- NOTE the spacing with the arrow when sending vs recieving. I think it will compile regardless of the spacing, but it may mess with the editor


======================================================================================


*/
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	//"strings"
)

func main() {

	
fmt.Println(trueRandomSeed())

	


}

func trueRandomSeed() int64{

	//Random.org true random number generator. Change the "len" number in the url to change the length
	url := "https://www.random.org/strings/?num=1&len=5&digits=on&unique=on&format=plain&rnd=new" 
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//change the byte array to a string
	s := string(html)

	//remove newline character
	s = s[:len(s)-1]

	//convert string to int64
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
    	panic(err)
	}

	return x;


}
