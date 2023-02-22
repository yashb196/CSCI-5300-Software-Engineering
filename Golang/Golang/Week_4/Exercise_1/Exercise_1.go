// Create an array which holds 5 values of TYPE int
// Assign values to each index position
// Range over the array and print the values out
// use format printing
// print out the type of array
package main

import (
	"fmt"
)

func main() {
	x := [5]int{42, 43, 44, 45, 46}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
