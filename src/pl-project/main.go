package main

// https://golang.org/doc/code.html
// http://www.newthinktank.com/2015/02/go-programming-tutorial/

import "fmt"
//import "customer.go"

// http://stackoverflow.com/questions/15049903/how-to-use-custom-packages-in-golang
// import "/pl-project/storeLibs" // Not working? according to above it should

type customer struct {
	id                 int
	serviceDurationSec int
}



func main() {
	//fmt.Println(customer{0, 5})
	fmt.Println("Test")
}
