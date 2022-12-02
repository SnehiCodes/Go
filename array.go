package main

import "fmt"

func main() {

	// var variable_name =[size]variable_type{}
	var bookings = [50]string{"Rahul", "Snehi", "Tuffy"}

	var friends = [50]string{}
	friends[0] = "Coco"
	friends[1] = "Puffy"
	friends[2] = "Limey"
	friends[3] = "Mukky"
	friends[04] = "Bunny"

	fmt.Printf("My all friends are : %v \n", friends)
	fmt.Printf("I want %v \n", bookings[0])

	fmt.Printf("Array type :%T \n", friends)

	fmt.Printf("Array length :%v \n", len(friends))
}
