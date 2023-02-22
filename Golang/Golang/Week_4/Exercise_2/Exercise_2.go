// using a composite literal :
// create a slice of type int
// assign 10 values
// range over the slice and print the values out
// use format printing
// print out the type of array
package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
