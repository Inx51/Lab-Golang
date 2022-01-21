package main

import "fmt"

func main() {

	fmt.Println("Enter a number: ")
	var number int
	fmt.Scanln(&number)

	if number > 0 {
		fmt.Println("Number is greater than zero")
	} else if number == 0 {
		fmt.Println("Number is zero")
	} else { // <- notice the { and then else
		fmt.Println("Number is lower than zero")
	}
}
