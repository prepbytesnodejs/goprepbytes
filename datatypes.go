package main

import "fmt"

/**
int
float
complex
string
bool
byte
rune
**/

// x64

func dataTypeGo() {

	var age int
	age = 9

	var floatValue float32 //declaration + initialization

	var complexNumber complex64

	var name string

	var iamGoingtoMArket bool

	var byteData byte

	var chars rune

	name = "joe"

	complexNumber = 2 + 4i

	floatValue = 20.2
	iamGoingtoMArket = true

	byteData = 2
	chars = 'a'

	fmt.Println(age)
	fmt.Println(floatValue)
	fmt.Println(complexNumber)
	fmt.Println(name)
	fmt.Println(iamGoingtoMArket)
	fmt.Println(byteData)
	fmt.Println(chars)

	//c := complex(23, 31)
	realPart := real(complexNumber) // gets real part
	imagPart := imag(complexNumber)

	fmt.Println(realPart)
	fmt.Println(imagPart)

	// Printing out the real part of a complex number in go

}
