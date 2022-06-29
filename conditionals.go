package main

import "fmt"

func checkCondition(num int) string {

	if num%2 == 0 && num > 0 {
		return "its even"

	} else if num%2 != 0 && num > 0 {
		return "its odd"

	} else {

		return "dnt pass negative number"
	}

}

// loops in go ---->
func simpleLoop() {
	//var i int

	for i := 0; i < 10; i++ {

		fmt.Println(i)

	}

}

//  While loop with for keyword

func whileLoop() {

	i := 0
	for i < 10 {
		fmt.Println(i * 9)
		i++
	}

}

// forEach loop

func forEach() {

	names := []string{
		"hershita",
		"junaid",
	}

	for i, s := range names {
		fmt.Println(i, s)

	}

}

func infiniteLoop() {

	for {
		fmt.Println("infinity")
	}

}
