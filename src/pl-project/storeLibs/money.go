package storeLibs

import "strconv"

type Money struct {
	Dollars int // It's a bad idea to store cost as a double because of floating point error.
	Pennies int // so we're storing the amount of dollars and amount of pennies seperately.
}

func (money *Money) ToString() string {
	return "$" + strconv.Itoa(money.Dollars) + "." + strconv.Itoa(money.Pennies)
}
