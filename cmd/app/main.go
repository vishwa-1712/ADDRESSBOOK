package main

import (
	"addressbook/config"
	"addressbook/internal/app"
	"fmt"
)

func main() {

	fmt.Println("WELCOME TO ADDRESS BOOK")

	config := config.LoadConfig()
	app.Run(config)
}
