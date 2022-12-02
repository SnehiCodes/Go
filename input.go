package main

import "fmt"

func main() {
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTickets int
	fmt.Scan(&userName)

	fmt.Println("How many tickets do you want ?")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v  booked %v tickets. \n", userName, userTickets)

}
