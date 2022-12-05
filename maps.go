// Maps

package main

import "fmt"

func main() {

	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	val, ok := mp["apple"]
	fmt.Println(val, ok)
	//mp := make(map[string]int)
	//fmt.Println(mp["apple"])
	//mp["apple"] =900
	//delete(mp,"apple")
	fmt.Println(mp)
	fmt.Println(len(mp))
}
