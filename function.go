package main

import (
	"fmt"
)

// function decalaration/function definition

func sum(a int, b int) { // parameter
	add := a + b

	fmt.Println(add)

}

// go allows you to return multiple values from a function

func subtract(a int, b int) int { // parameter
	minus := a - b

	return minus

}
