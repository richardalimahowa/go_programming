package main

import "fmt"

var pl = fmt.Println

func main() {
	// for initialization; condition; poststatement { body}
	/*for x := 5; x >= 1; x-- {
		pl(x)

	}
	// while loop
	fx := 1
	for fx <= 5 {
		pl(fx)
		fx++
	}*/

	// arrays
	/*aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}*/
	// do while loop
	xVal := 1
	for true {
		if xVal == 5 {
			break
		}
		pl(xVal)
		xVal++
	}

}
