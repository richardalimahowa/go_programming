package main

import (
	"bufio"
	"fmt"
	f "fmt"
	"log"
	"os"
)

var pl = f.Println

func main() {
	pl("hello Go")

	fmt.Print("what is your name ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)

	} else {
		log.Fatal(err)
	}

}
