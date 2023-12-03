package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// if conditional
	// conditional operators : > < >= <= == !=
	// logical : && || !
	/*iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")

	} else if iAge >= 65 {
		pl("Important Birthday")

	} else {
		pl("Not an Important Birthday")
	}
	*/
	//pl("!true = ", !true)
	//fmt.Printf("%s %d %c %f %t %o %x\n", "stuff", 1, 'A', 3.14, true, 1, 1)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f", 3.141592)
	fmt.Printf("%9.f", 3.141592)
	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)

}
