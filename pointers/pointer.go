package pointers

import "fmt"

func Pointers() {

	// Pointers are variables that hold address of another variable
	a := 9

	b := &a // b is pointer to a

	// * --->its a dereferencing operator

	fmt.Println(a)
	fmt.Println(b) ///Printing the pointer

	fmt.Println(*b) // derefrencing the pointer

}

// Passing pointer to a function as an argument
// go is always pass by value ,go is not pass by reference

func TakePointer(a *int) { // this function will take pointer to an integer varianble

	*a = 88

}

// Getting return value as pointer from function

func ReturnsPointer() *string {

	var g *string
	var k string = "hello there"

	g = &k

	return g

}

//Pointer to slice  in function arguments

func PointerToSlice(ptr *[]int) {

	fmt.Println(ptr)

	// append method is used to add values to an existing slice

}
