package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	var selection int
	fmt.Println("Print a pithy saying!")
	fmt.Println("What would you like to hear?")
	fmt.Println("Pick from:1,2,or 3")
	fmt.Scanln(&selection)
	if selection == 1 {
		fmt.Println(quote.Go())
		return
	} else if selection == 2 {
		fmt.Println(quote.Glass())
		return
	} else if selection == 3 {
		fmt.Println(quote.Opt())
	} else {
		fmt.Println("Wrong selection no loops fool")
		fmt.Println("Run: > go run pithyello.go again")
	}
}
