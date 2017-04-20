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
	- Channels can be closed with close(), but it is not necessary. If you try to pull from a closed channel, you get a zero value
	- Channels are bi directional by default, but can be restricted
		var recvOnly <-chan int
		var sendOnly chan<- int
	- Channels have a control structure called "select", which functions similarly to a switch statement. Case statements must be channels

		select{	
			case x := <-ch0:
				fmt.Println(x)
			case y := <-ch1:
				fmt.Println(y)
			case z := <-ch2:
				fmt.Println(x)
			default:
				fmt.Println("default")	
		}

	- Basically it checks all channels in the statement to see if any are ready to recieve
	- If none are ready, it defaults. If there is no default, it blocks. If more than 1 is ready it "fairly" randomly selects one to recieve

======================================================================================


*/
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	//"time"
	"math/rand"
	//"strings"
)

func main() {

	var wg sync.WaitGroup
	s := rand.NewSource(trueRandomSeed())
	r := rand.New(s)
	var ch0 = make(chan int)

	fmt.Println(r.Intn(100))

	wg.Add(1)
	go func(){
		fmt.Println("start")
		select{
			case ch0 <- 5:
				fmt.Println("sent 5")
			default: 
				fmt.Println("default")
				fmt.Println("default2")
		}

		wg.Done()
	}()


	y := <- ch0

	fmt.Println(y)
	

	wg.Wait()
}

func trueRandomSeed() int64{

	//Random.org true random number generator. Change the "len" number in the url to change the length
	url := "https://www.random.org/strings/?num=1&len=18&digits=on&unique=on&format=plain&rnd=new" 
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
