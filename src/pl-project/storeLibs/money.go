package storeLibs

import "strconv"

type Money struct {
	Dollars int // It's a bad idea to store cost as a double because of floating point error.
	Pennies int // so we're storing the amount of dollars and amount of pennies seperately.
}

func (money *Money) Add(other *Money) {
	money.Dollars += other.Dollars
	money.Pennies += other.Pennies

	money.Dollars += int(money.Pennies / 100)
	money.Pennies %= 100
}

func (money *Money) ToString() string {
	var pennies string

	if money.Pennies < 10 {
		pennies = "0" + strconv.Itoa(money.Pennies)
	} else {
		pennies = strconv.Itoa(money.Pennies)
	}

	return "$" + strconv.Itoa(money.Dollars) + "." + pennies
}
