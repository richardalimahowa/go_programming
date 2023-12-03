package main

import "fmt"

func main() {
	//fmt.Println(42)
	//fmt.Printf("%d - %b \n", 42, 42)
	//fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	//fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)
	//for i := 60; i < 122; i++ {
	//fmt.Printf("%d \t %b \t %x \t %q  \n", i, i, i, i)
	//}
	/*x := 13 % 3
	fmt.Println(x)

	for i := 1; i < 70; i++ {
		if i%2 == 1 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}

	}*/
	/*switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("you have no friends")
	}*/
	/*switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("wassup Jenny")
	case "Marcus":
		fmt.Println("wassup Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}*/

	/*switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("wassup julian / Sushant")
	}*/

	/*myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("wassup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched")
	}*/
	// if else statements
	/*if true {
		fmt.Println("This ran")
	}
	if false {
		fmt.Println("This did not run")
	}*/

	/*b := true

	if food := "chocolate"; b {
		fmt.Println(food)
	}*/
	/*for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}

	}*/

	fmt.Println(greet("Jane", "Doe"))

}
func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
