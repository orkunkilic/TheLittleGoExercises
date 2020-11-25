package shopping

import (
	"./db" //This can be $GOPATH/src/shopping/db
)

func PriceCheck(itemId int) (float64, bool) { //if it is priceCheck, than we won't access it in main.go
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
