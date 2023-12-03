package main

import (
	"fmt"
)

/*
	func getSum(x int, y int) int {
		return x + y
	}
*/
/*func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You cant divide by zero")
	} else {
		return x / y, nil
	}
}*/
/*func getTwo(x int) (f1 int, f2 int) {
	return x + 1, x + 2
}*/
/*func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}*/

/*
	func getArrSum(arr []int) int {
		sum := 0
		for _, num := range arr {
			sum += num
		}
		return sum
	}
*/
func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func main() {
	/*for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fX := 1 // WHILE LOOP
	for fX <= 5 {
		fmt.Println(fX)
		fX++

	}*/

	/*xVal := 1
	for true {
		if xVal == 5 {
			break
		}
		fmt.Println(xVal)
		xVal++
	}*/
	/*reader := bufio.NewReader(os.Stdin)
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("Guess a number between 1 and 50:")
		fmt.Println("Random Number is :", randNum)
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
			fmt.Println("Guess Something Lower")

		} else if iGuess < randNum {
			fmt.Println("Guess something Higher")

		} else {
			fmt.Println("you guessed it")
			break
		}
	}*/

	// Guess a number between 0 and 50
	// Random Number is : 25

	/*sV1 := " A word"
	// Escape sequences : \n \t \
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	fmt.Println(sV2)
	fmt.Println("Length :", len(sV2))
	fmt.Println("contains Another :", strings.Contains(sV2, "Another"))
	fmt.Println("o index", strings.Index(sV2, "o"))
	fmt.Println("Replace :", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\nSome words \n"
	sV3 = strings.TrimSpace(sV3)
	fmt.Println("Split :", strings.Split("a-b-c-d", "-"))
	fmt.Println("Lower:", strings.ToLower(sV2))
	fmt.Println("upper:", strings.ToUpper(sV2))
	fmt.Println("Prefix:", strings.HasPrefix("tacocat", "taco"))
	fmt.Println("suffix:", strings.HasSuffix("tacocat", "cat"))*/

	// runes are characters
	/*rStr := "abcdefg"
	fmt.Println("Rune Count:", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}*/
	// Time

	/*now := time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day())
	fmt.Println(now.Hour(), now.Minute(), now.Second())*/

	// changing location
	/*loc, err := time.LoadLocation("Uganda/Kampala")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("Time in New York %S\n", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Shanghai")
	fmt.Printf("Time in shanghai %s\n", now.In(loc))

	locEST, _ := time.LoadLocation("EST")
	locUTC, _ := time.LoadLocation("UTC")
	locMST, _ := time.LoadLocation("MST")
	fmt.Printf("EST : %s\n", now.In(locEST))
	fmt.Printf("UTC : %s\n", now.In(locUTC))
	fmt.Printf("MST : %s\n", now.In(locMST))*/

	// calculating time
	/*birthDate := time.Date(1974, time.December, 21, 11, 30, 10, 0, time.Local)
	diff := now.Sub(birthDate)
	fmt.Printf("Days Alive : %d days \n", int(diff.Hours()/24))
	fmt.Printf("Hours Alive : %d hours \n", int(diff.Hours()))*/

	// arrays
	/*var arr1 [5]int
	arr1[0] = 1
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("index 0 :", arr2[0])
	fmt.Println("Array Length :", len(arr2))
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arr3[i][j])
		}
	}*/
	// slicing a string into a rune
	/*astr11 := "abcde"
	rArr := []rune(astr11)
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}
	// byte array to string
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	fmt.Println("Iam a string :", bStr)*/
	// slices
	// var name []datatype
	/*sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	fmt.Println("slice size", len(sl1))
	for i := 0; i < len(sl1); i++ {
		fmt.Println(sl1[i])

	}
	for _, x := range sl1 {
		fmt.Println(x)
	}
	// slice literal
	sl2 := []int{12, 21, 1947}
	fmt.Println(sl2)

	sArr := [5]int{1, 2, 3, 4, 5}
	//sl3 := sArr[0:2]
	//sl3 := sArr[:3]
	//fmt.Println("this is first three", sl3)
	//sl3 := sArr[2:]
	sl3 := sArr[0:2]
	fmt.Println("last three", sl3)

	sArr[0] = 10
	fmt.Println("sl3 :", sl3)

	// appending
	sl3 = append(sl3, 12)
	fmt.Println(sl3)
	fmt.Println("sArr :", sArr)

	// empty slice
	sl4 := make([]string, 6)
	fmt.Println("sl4 :", sl4)
	fmt.Println("sl4[0] :", sl4[0])*/

	// functions
	//fmt.Println(getSum(1, 2))
	/*ans, err := getQuotient(5, 0)
	if err == nil {
		fmt.Printf("%f / %f = %f", 5.0, 0, ans)
	} else {
		fmt.Println(err)
	}*/
	/*f1, f2 := getTwo(5)
	fmt.Printf("%d %d\n", f1, f2)*/
	//fmt.Println("unknown sum:", getSum2(1, 2, 3, 4))
	/*vArr := []int{1, 2, 3, 4}
	fmt.Println("Array Sum:", getArrSum(vArr))*/
	/*f3 := changeVal(5)
	fmt.Println("f3 before func :", f3)
	/*changeVal(f3)
	fmt.Println("f3 after func:", f3)*/
	// hang man exercise

}
