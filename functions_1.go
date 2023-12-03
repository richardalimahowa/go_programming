package main

import (
	"fmt"
)

/*func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}*/

/*func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}*/

// arrays
func main() {
	//var x [58]string
	/*fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	x[42] = 777
	fmt.Println(x[42])*/

	/*for i := 65; i <= 122; i++ {
		x[i-65] = string(i)

	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])*/

	/*var x [256]byte

	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i < 256; i++ {
		x[i] = byte(i)

	}
	for i, v := range x {
		fmt.Print("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}*/
	/*var x [256]string

	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i < 256; i++ {
		x[i] = string(i)

	}
	for i, v := range x {
		fmt.Print("%v - %T - %v\n", v, v, []byte(v))
		if i > 50 {
			break
		}
	}*/

	/*mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)*/
	/*mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  //slicing a slice
	fmt.Println(mySlice[2])    // index access; accessing by index
	fmt.Println("myString"[2]) // index access; accessing by index
	*/
	/*student := []string{} // shorthand
	students := [][]string{}
	student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)*/

	/*var student []string // var method
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)*/

	/*student := make([]string, 35) // using make
	students := make([][]string, 35)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)*/

	/*mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7

	fmt.Println(mySlice[0])
	mySlice[0]++
	mySlice[0] = mySlice[0] + 1
	fmt.Println(mySlice[0])*/

	/*var nonce [24]byte // code didnot run
	fmt.Println(nonce)
	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)*/

	//var myGreeting = make(map[string]string)// option 1
	/*var myGreeting = map[string]string{} // otion 2

	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	fmt.Println(myGreeting)*/

	myGreeting := map[string]string{
		"0": "Good morning",
		"1": "Bonjour!",
	}
	//myGreeting["Harleen"] = "Howdy"
	delete(myGreeting, "0")
	fmt.Println(myGreeting)
}
