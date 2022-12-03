package main

import "fmt"

func main() {
	x := 5
	y := 6.5

	fmt.Printf("%t \n", float64(x)+1.5 != y)

	j := "b"
	k := "A"
	val := j > k
	fmt.Printf("%t", val)
}
