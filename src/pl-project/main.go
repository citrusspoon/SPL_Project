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

var currentId = 0
const MINSERVICEDURATION int = 60
const MAXSERVICEDURATION int = 600


func main() {

	fmt.Println(createCustomer())
	fmt.Println(createCustomer())
	fmt.Println(createCustomer())
	fmt.Println(createCustomer())


}

func createCustomer() customer{

	s1 := rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000000)))
 	r1 := rand.New(s1)

	currentId++ //for some reason this can't be passed as an argument
	return customer{currentId, MINSERVICEDURATION + r1.Intn(MAXSERVICEDURATION - MINSERVICEDURATION)}
}
