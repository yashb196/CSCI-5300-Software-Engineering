// print out years you have been alive
package main

import (
	"fmt"
)

func main() {
	bd := 1996

	for bd <= 2022 {
		if bd > 2022 {
			break
		}
		fmt.Println(bd)
		bd++

	}
}
