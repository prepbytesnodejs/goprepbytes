package structs

import "fmt"

// struct is a composite  user defined data type
// in which you can multiple different types together

//type MyInt int

// structs in go
type Student struct {
	Name            string
	Age             int
	IsComingToClass bool
	isComingToClass bool
}

// Go has concept of receivers
// receiver pointer
// receive value
// method
// a method can be a function but a function can not be method

func GetName() { // this cant be directly invoked without instantiating the struct

	fmt.Println("I am a structs method")

}

func (s *Student) GetName() { // this cant be directly invoked without instantiating the struct

	fmt.Println("I am a structs method")

}

func (s *Student) AssignValues(name string) { // this cant be directly invoked without instantiating the struct
	s.Name = name
	//	fmt.Println("I am a structs method")

}

func (s Student) Somemethod(name string) { // this cant be directly invoked without instantiating the struct
	s.Name = name
	//	fmt.Println("I am a structs method")

}

// this can be a revceiver value

/*
func Mytype(a int) bool {

	var k MyInt
	k = 9

	if k==a{  // not possible

	}

}
*/
