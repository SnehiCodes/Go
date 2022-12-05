package main

import "fmt"

func test() {
	fmt.Println("Hello !")
}
func main() {
	x := test
	x()
	main1()
	test3 := func(x int) int {
		return x * 7
	}
	test3(7)
}
func main1() {
	test1 := func() {
		fmt.Println("hello !")
	}
	test1()

}
func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}
