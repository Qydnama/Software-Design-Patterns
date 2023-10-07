package main

import "fmt"

func main() {
	lays, _ := getChips("lays")
	hrustim, _ := getChips("hrustim")

	printDetails(lays)
	printDetails(hrustim)
}

func printDetails(c IChips) {
	fmt.Printf("Chips: %s", c.getName())
	fmt.Println()
	fmt.Printf("Rating: %d", c.getRating())
	fmt.Println()
}
