package main

import (
	"fmt"
	"strings"
)

func main() {
	//Array (value type)
	arr := [3]string{"Hello", "world"}
	arr[1] = "array"
	fmt.Println(arr[0] + " " + arr[1] + "!")

	//Slice (ref type)
	slice := []string{"Hello", "world"}
	slice[1] = "from"
	slice = append(slice, "slice!")
	fmt.Println(strings.Join(slice, " "))

	//Range selection
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rangeSkip := numbers[2:]
	fmt.Println(rangeSkip)
	rangeTake := numbers[:3]
	fmt.Println(rangeTake)
	rangeSkipTake := numbers[2:5]
	fmt.Println(rangeSkipTake)
}
