package main

import (
	"fmt"
	"goprepbytes/goconcurrency"
	_ "goprepbytes/package2/package2" // this syntax is used to import pacakges for side effects only
	_ "goprepbytes/pointers"
	_ "goprepbytes/rangego"
	_ "goprepbytes/structs"
	"sync"
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
	/*
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
	/*

		// Taking user input in go

		var user string

		fmt.Println(&user)

		fmt.Println("Enter your name")

		//fmt.Scanf("")

		fmt.Println("Hello", user)
	*/
	//fmt.Println(takeInputWithSpaces())

	//dataTypeGo()
	//a := checkCondition(3)

	//fmt.Println(a)
	//simpleLoop()

	//whileLoop()

	//forEach()

	//infiniteLoop()
	//arrays()

	//sliceS := []int{9, 9}

	//array := [2]int{88, 77}

	//fmt.Println(array)

	//fmt.Println(sliceS)

	//modifySlice(sliceS)

	//fmt.Println(sliceS)
	//modifyArray(array)
	//fmt.Println(array)
	//typeConversion()
	//mapMethod()

	//var i int = 1
	//var j float32 = 1
	//k := i == j

	//package2.Foo()

	//pointers.Pointers()

	//a := 9
	//pointers.TakePointer(&a)

	//fmt.Println(a)

	//fmt.Println(*pointers.ReturnsPointer())

	// one is make and one is new
	// new keyword is used to create slices ,structs  // new keyword returns a pointer
	// make keyword is used to create maps and chan,slice

	//slc2 := make([]int, 10)

	//slcPointer := new([]int)

	//pointers.PointerToSlice(slcPointer)

	//We have to instantiate the struct first
	/*
		sobj := structs.Student{
			Name:            "John",
			Age:             20,
			IsComingToClass: true,
		} //instantiating the struct

		fmt.Println(sobj.Name, sobj.Age, sobj.IsComingToClass)
		//structs.GetName() // calling a function
		//sobj.GetName()    //calling a method

		emptyStruct := structs.Student{
			Name: "Johnny Walker",
		}

		emptyStruct.AssignValues("John cena")

		fmt.Println(emptyStruct.Name)
	*/

	// value receiver
	/*
			emptyStruct := structs.Student{}

			emptyStruct.AssignValues("John cena")

			fmt.Println(emptyStruct) ///

			struct2 := structs.Student{}

			struct2.Somemethod("Random name")

			fmt.Println(struct2) ///


		slc := []rangego.Student{} // A slice of custom type

		for i := 0; i < 4; i++ {
			slc = append(slc, rangego.SliceOfStruct("john", 88))

		}


		// Range operator // Its kind of enhanced for loop

		// forEach loop from other programming languages

		for i, val := range slc {
			fmt.Println(i, val)

		}

		//iterate over a map
		/*
			mp := make(map[string]string)
			mp["name"] = "John"
			mp["Location"] = "Pune"
			mp["Pincode"] = "110001"
			mp["Country"] = "India"

			for key, val := range mp {
				fmt.Println(key, val)
			}
	*/

	var waitGroup sync.WaitGroup //it groups all over go routines and
	// we can wait inside the main routine till all our all the go routines
	// have finished execution

	for i := 0; i < 5; i++ {
		go goconcurrency.GoRoutine(i, &waitGroup)
		//go goconcurrency.GoRoutine2(i)
	}

	//time.Sleep(time.Second)
	waitGroup.Wait()
	fmt.Println("Main exited")

}
