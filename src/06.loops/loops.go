package main

import "fmt"

func main() {

	//Normal for-loop
	for i := 0; i <= 10; i++ {
		fmt.Print(fmt.Sprint(i) + " ")
	}

	//While
	j := 0
	for j <= 10 {
		fmt.Print(fmt.Sprint(j) + " ")
		j++
	}

	//For-each
	k := []string{"Apple", "Banana", "Pear"}
	for _, v := range k { // <- _ discard the index
		fmt.Print(v + " ")
	}
}
