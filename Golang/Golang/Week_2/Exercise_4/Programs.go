// Write a program that
// assigns an int to a variable
// Prints that int in decimal, binary and hex
// Shifts the bits of that int over 1 position to the left, and assigns that to a variable
// Prints the variable in decimal, binary and hex
package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}
