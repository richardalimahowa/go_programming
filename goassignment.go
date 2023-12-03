package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	// Receive customer data(Their age)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("what is your age : ")
	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)

	}

	// google how to trim whitespace from input
	age = strings.TrimSpace(age)
	iAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}
	// implementing logic
	if iAge < 5 {
		fmt.Println("Too young for school")

	} else if iAge == 5 {
		fmt.Println("Go to Kindergarden")
	} else if (iAge > 5) && (iAge <= 17) {
		grade := iAge - 5
		fmt.Printf("Go to grade %d\n", grade)

	} else {
		fmt.Println("Go to college")
	}

}
