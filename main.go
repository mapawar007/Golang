package main

import "fmt"

func main() {
	fmt.Println("We are learning array in golang")

	var name [5]string

	name[1] = "Manish"
	name[4] = "Pawar"
	fmt.Println("names of person is :", name)

	var number = [8]int{1, 2, 3, 4, 5}
	fmt.Println("Number is :", number)

	fmt.Println("Length of numbers is :", len(number))
	fmt.Println("value of name at 2 index is :", name[1])
}
