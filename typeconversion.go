package main

import "fmt"

func typeConversion() {

	var x int = 8
	var y int = 9

	var z float32

	z = float32(x) * float32(y)

	fmt.Printf("Float number %f ", z)
}
