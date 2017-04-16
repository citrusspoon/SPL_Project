package storeLibs

import (
	"math/rand"
	"strconv"
	"time"
)

var currentID = 0

type Customer struct {
	ID                 int
	ServiceDurationSec int
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
	return "ID: " + strconv.Itoa(customer.ID) + ", Service Time (seconds): " + strconv.Itoa(customer.ServiceDurationSec)
}

func MakeCustomer() *Customer {

	/*
		This is necessary to reset the seed that the random function uses. Otherwise the numbers generated would be the same every time
		as Go doesn't reset the seed on its own. Just using the time isn't enough because it runs too fast, so I threw in an extra
		random number as well. It seems to be somewhat random.
	*/
	s1 := rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000000)))
	r1 := rand.New(s1)

	const MINSERVICEDURATION int = 60
	const MAXSERVICEDURATION int = 600

	currentID++

	return &Customer{currentID, MINSERVICEDURATION + r1.Intn(MAXSERVICEDURATION-MINSERVICEDURATION), nil}
}

// 1% chance someone is added per second
func IsCustomerAdded() bool {
	s := rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000000)))
	r := rand.New(s)

	return r.Intn(100) == 0
}
