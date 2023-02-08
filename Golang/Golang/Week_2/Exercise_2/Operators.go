//Using the following operators, write expressions and assign their values to variables:
// g. ==
// h. <=
// i. >=
// j. !=
// k. <
// l. >

package main

import (
	"fmt"
)

func main() {
	a := (12 == 12)
	b := (12 <= 13)
	c := (12 >= 13)
	d := (12 != 13)
	e := (12 > 13)
	f := (12 < 13)
	fmt.Println(a, b, c, d, e, f)
}
