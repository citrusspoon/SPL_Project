package storeLibs

import (
	"math/rand"
	"time"
)

type Item struct {
	ID      int
	Name    string
	Price   *Money
	Quanity int
}

// here at harris lion we've got it all.  except for what we don't have
var StoreItems = []*Item{
	&Item{0, "Bread", &Money{2, 99}, 100},
	&Item{1, "Chicken Breast", &Money{3, 99}, 52},
	&Item{2, "Fish", &Money{3, 24}, 55},
	&Item{3, "Pasta", &Money{1, 22}, 44},
	&Item{4, "Rice", &Money{1, 99}, 43},
	&Item{5, "Cereal", &Money{2, 99}, 41},
	&Item{6, "Cheese", &Money{0, 99}, 94},
	&Item{7, "Soup", &Money{0, 49}, 100},
}

// attempts to buy an item.  If it's not in stock it returns false
func buyItem(index int) bool {
	if StoreItems[index].Quanity <= 0 {
		return false
	}

	StoreItems[index].Quanity--
	return true
}

func Price(items []*Item) *Money {
	money := &Money{0, 0}

	for _, item := range items {
		money.Add(item.Price)
	}

	return money
}

func GetItems() []*Item {
	var items []*Item

	s := rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000000)))
	r := rand.New(s)

	const maxPossibleItems = 15
	const minPossibleItems = 1

	for item := 0; item < r.Intn(maxPossibleItems-minPossibleItems)+minPossibleItems; item++ {
		index := r.Intn(len(StoreItems))

		if buyItem(index) {
			items = append(items, StoreItems[index])
		}
	}

	return items
}
