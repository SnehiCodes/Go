package main

import "fmt"

func main() {

	answer := 10
	switch answer {
	case 10:
		fmt.Println("1. one")

	case 2:
		fmt.Println("two")

	default:
		fmt.Println("not a case")
	}
}
