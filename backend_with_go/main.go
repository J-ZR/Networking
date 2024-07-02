package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello, world")

	/*

		GO variables:
		int - stores integers
		float32 - stores floating point numbers with decimals
		string - stores text
		bool - stores values with two states: true or false

	*/

	// declaring a variable
	var my_var string = "hello" // type is string
	var second_var = "nice"     // type is inferred
	x := 2                      // type is inferred

	// variable declaration without initial value

	var a string
	var b int
	var c bool

	// variable assignment after declaration
	var student1 string
	student1 = "jonn"
	student2 := "anna"

	// outputs
	fmt.Println(student1)
	fmt.Println(student2)

	fmt.Println(my_var)
	fmt.Println(second_var)
	fmt.Println(x)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// variable decleration in a block
	var (
		num1      int
		num2      int    = 1
		my_string string = "hello"
		//d := 29.8
	)

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(my_string)
	//fmt.Println(d)

}
