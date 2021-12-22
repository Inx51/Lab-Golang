package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	//Declar multiple variables
	var declaredVariableOne, declaredVariableTwo string

	//Declare and init single var
	var declaredAndInitVariableOne = "Hello, world"

	//Shorter version to declare and init without explicitly setting type
	declaredAndInitVariableTwo := 123
	declaredAndInitVariableThree := true

	fmt.Println("Declared variable: " + declaredVariableOne)
	fmt.Println("Declared variable: " + declaredVariableTwo)
	fmt.Println("Declared and init variable: " + declaredAndInitVariableOne)
	fmt.Println("Declared and init variable: " + fmt.Sprint(declaredAndInitVariableTwo))
	fmt.Println("Declared and init variable: " + fmt.Sprint(declaredAndInitVariableThree))
}
