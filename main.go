package main

import (
	"fmt"
)

/*
	func main() {
		var name string = "Timothy"
		var age int = 35
		var married bool = true

		fmt.Println(name, age, married)
	}
*/
/*func main() {
	var name string = "Timothy"
	var weight int = 225
	var height float64 = 5.9

	//fmt.Println("welcome to 2024 christmas day")
	//fmt.Println("my name is ", name, " i am ", weight, " pounds and my height is ", height)
	fmt.Printf(" welcome to the 2024 christmas day, %v! Please confirm you are %2f height  and %d in weight?", name, height, weight)
}*/

func main() {
	var name string
	var age int
	fmt.Print("What is your name and age? ")
	fmt.Scanf("%s, %d", &name, &age)
	fmt.Println("thank you", name, age)
}
