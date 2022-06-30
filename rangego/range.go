package rangego

type Student struct {
	Name string
	Age  int
}

func SliceOfStruct(name string, Age int) Student {

	objStruct := Student{name, Age}

	return objStruct

}
