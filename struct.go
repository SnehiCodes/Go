package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	episodes   []string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",

		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])
}
