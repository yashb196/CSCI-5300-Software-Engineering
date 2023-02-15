// print every year from your birthdate to current year
package main

import (
	"fmt"
)

func main() {
	bd := 1996
	for bd <= 2022 {
		fmt.Println(bd)
		bd++
	}
}
