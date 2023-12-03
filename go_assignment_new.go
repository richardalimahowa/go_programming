package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {
	//guess a number
	reader := bufio.NewReader(os.Stdin)
	seedSecs := time.Now().Unix() // creating a random number
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("guess a number between 1 and 50:")
		pl("Random Number is :", randNum)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Guess Something Lower")
		} else if iGuess < randNum {
			pl("Guess something higher")

		} else {
			pl("you guessed it")
			break
		}

	}

}
