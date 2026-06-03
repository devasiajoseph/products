package main

import (
	"fmt"

	"github.com/adoniaromal/products/db"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		fmt.Println("error connecting db")
		fmt.Println(err)
		return
	}

	fmt.Println(db)

}
