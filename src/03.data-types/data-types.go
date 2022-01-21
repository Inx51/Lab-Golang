package main

import "fmt"

func main() {
	s := "string"
	i := 1
	f := 10.2
	b := true
	by := 0x01

	fmt.Println("String: " + s)
	fmt.Println("Integer: " + fmt.Sprint(i))
	fmt.Println("String: " + fmt.Sprint(f))
	fmt.Println("String: " + fmt.Sprint(b))
	fmt.Println("String: " + fmt.Sprint(by))
}
