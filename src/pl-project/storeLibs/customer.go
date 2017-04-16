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
	return "ID: " + strconv.Itoa(customer.ID) + ", Service Time (seconds): " + strconv.Itoa(customer.ServiceDurationSec)
}

func MakeCustomer() *Customer {
	currentID++

	const secondsPerItem = 5
	items := GetItems()

	return &Customer{currentID, secondsPerItem * len(items), items, nil}
}

// 1% chance someone is added per second
func IsCustomerAdded() bool {
	s := rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000000)))
	r := rand.New(s)

	return r.Intn(100) == 0
}
