package storeLibs

import (
	"math/rand"
	//"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
	"time"
)

var currentID = 0

type Customer struct {
	ID                 int
	ServiceDurationSec int
	Items              []*Item
	NextCustomer       *Customer
}

func (customer *Customer) DecServiceTime() {
	customer.ServiceDurationSec--
}

// returns true if the customer has finished being serviced, false otherwise
func (customer *Customer) Service() bool {
	customer.DecServiceTime()
	return customer.HasBeenServiced()
}

func (customer *Customer) HasBeenServiced() bool {
	return customer.ServiceDurationSec <= 0
}

func (customer *Customer) ToString() string {
	return "ID: " + strconv.Itoa(customer.ID) + ", Service Time (seconds): " + strconv.Itoa(customer.ServiceDurationSec) + ", Price of items: " + Price(customer.Items).ToString()
}

func MakeCustomer() *Customer {
	currentID++

	const secondsPerItem = 5
	items := GetItems()

	return &Customer{currentID, secondsPerItem * len(items), items, nil}
}

// 50% chance someone is added per second
func IsCustomerAdded() bool {
	s := rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000000)))
	//s := rand.NewSource(trueRandomSeed())
	r := rand.New(s)

	return r.Intn(100) == 0
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