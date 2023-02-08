// Create a variable to type string using a raw string literal. Print It
package main

import (
	"fmt"
)

func main() {
	a := `Here is something 
	as
	a
	raw string
	literal string 
	literal
	"you see"
	another things`
	fmt.Println(a)
}
