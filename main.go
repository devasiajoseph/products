package main

import (
	"fmt"

	"github.com/adoniaromal/products/db"
	"github.com/adoniaromal/products/pkg/items"
)

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Println("error connecting db")
		fmt.Println(err)
		return
	}
	fmt.Println(db.DBcon)

	items.CreateTag("item-tag")

}
