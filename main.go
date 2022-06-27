package main

import (
	"fmt"
)

// function to return multiple values is go

func arithemetic(a int, b int) (int, int) {

	divide := a / b
	multiply := a * b

	return divide, multiply

}

func arithemeticmodulus(a int, b int) (int, int, int) {

	divide := a / b
	multiply := a * b
	modulus := a % b

	return divide, multiply, modulus

}

/**
 */
// Is the starting point of your application
func main() {
	/*
		var (
			a,
			b,
			c int
		)
		a, b, c = arithemeticmodulus(6, 2)
	*/

	a, b, c := arithemeticmodulus(6, 2)

	fmt.Println(a, b, c)
	/*

		sum(3, 4) // function call

		result := subtract(8, 9)

		d, m := arithemetic(10, 2) // it returns two values
		_, q, s := arithemeticmodulus(9, 2)

		fmt.Println(q, s)

		fmt.Println(d, m)

		fmt.Println(result)
		/*
			//fmt.Println("Hello Mom !!!!")

			// variable declaration
			var age int //only variable declaration
			age = 20

			fmt.Println(age)

			// shorthand variable decalaration
			var x int //delare

			x = 9 // assignments

			fmt.Println(x)

			y := "hello world" // decalaration and assignment at the same place
			fmt.Println(y)

			// const variable
			//var name string = "john"
			//name = "joe" // this is not possible because of const keyword

			//fmt.Println(name)
	*/

	// Taking user input in go

	var user string

	fmt.Println(&user)

	fmt.Println("Enter your name")

	//fmt.Scanf("")

	fmt.Println("Hello", user)

}
