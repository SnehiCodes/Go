// infinite for loop

package main

import "fmt"

func main() {

	for {
		var name string
		fmt.Println("What's your name chum ?")
		fmt.Scan(&name)
		fmt.Printf("Hello , %v .\n", name)
	}

}
