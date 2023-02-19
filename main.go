package main

import (
	"fmt"
	"strings"
	"test/helpers"
)

func main() {
	helpers.LoadRulesFromConfig()
	for {
		fmt.Print("Enter pin: ")
		var input string
		fmt.Scanln(&input)
		pin := strings.Trim(input, "\n")
		fmt.Printf("Valid: %+v\n", helpers.ValidatePin(pin))
	}
	// fmt.Printf("Valid: %+v\n", helpers.ValidatePin("153474"))
}
