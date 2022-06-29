package main

import "fmt"

func arrays() {

	arrayInt := [2]int{0, 9} // its an array

	arrayInt2 := [...]int{0, 9} // its an array

	sliceInt := []int{8, 9} // its a slice

	fmt.Println(arrayInt)
	fmt.Println(arrayInt2)

	fmt.Println(sliceInt)
}

func modifySlice(a []int) {
	a[1] = 99
}

func modifyArray(a [2]int) {
	a[1] = 890

}
