package main

import (
	"fmt"
	"lamvng/finance-tracker/db"
)

func main() {
	DB := db.Init()
	fmt.Println(DB)
}
