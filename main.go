package main

import (
	"fmt"

	"github.com/REIS0/scuz/database"
)

func main() {
	// TODO: testar banco de dados
	db, err := database.Database.InitDatabase("scuz.sqlite")
	if err != nil {
		fmt.Println("database error")
		return
	}

}
