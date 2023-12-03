package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	CV1 := 1.5
	CV2 := int(CV1)
	pl(CV2)

	CV3 := "5000000"
	CV4, err := strconv.Atoi(CV3)
	pl(CV4, err, reflect.TypeOf(CV4))

	CV5 := 50000000
	CV6 := strconv.Itoa(CV5)
	pl(CV6)

	CV7 := "3.14"
	if CV8, err := strconv.ParseFloat(CV7, 64); err ==
		nil {
		pl(CV8)

	}

	CV9 := fmt.Sprintf("%f", 3.14)
	pl(CV9)

}
