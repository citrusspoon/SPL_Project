package main

// https://golang.org/doc/code.html
// http://www.newthinktank.com/2015/02/go-programming-tutorial/

import ( 
	"fmt"
	"math/rand"
	"time"

)
//import "customer.go"

// http://stackoverflow.com/questions/15049903/how-to-use-custom-packages-in-golang
// import "/pl-project/storeLibs" // Not working? according to above it should

type customer struct {
	id                 int
	serviceDurationSec int
}

type register struct {
	id                 int
	queueSize		   int
	maxQueueSize	   int
	currentlyServicing bool
}

var currentId = 0
const MINSERVICEDURATION int = 60
const MAXSERVICEDURATION int = 600
const NUMOFREGISTERS int = 5


func main() {

	//var registers[NUMOFREGISTERS] register

	fmt.Println("test")


}

func createCustomer() customer{

	/*
		This is necessary to reset the seed that the random function uses. Otherwise the numbers generated would be the same every time 
		as Go doesn't reset the seed on its own. Just using the time isn't enough because it runs too fast, so I threw in an extra 
		random number as well. It seems to be somewhat random. 
	*/
	s1 := rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000000))) 
 	r1 := rand.New(s1)

	currentId++ //for some reason this can't be passed as an argument
	return customer{currentId, MINSERVICEDURATION + r1.Intn(MAXSERVICEDURATION - MINSERVICEDURATION)}
}
