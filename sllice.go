// slice is an abstraction of array

package main

import "fmt"

func main() {
	var booking []string
	// another way to write it is...
	// var booking =[]string{}
	// or
	// booking := []string{}

	booking = append(booking, "Muku")
	booking = append(booking, "Coolfie")

	fmt.Printf("I booked  : %v \n", booking)
}
