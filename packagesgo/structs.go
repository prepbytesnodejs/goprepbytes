package packagesgo

import (
	"goprepbytes/package2"
)

type Student struct {
	Name string
	Age  int
}

//if you want to export a function in go make sure you are declaraing
// it with CapitalLetter

func Iamfromdifferentpackage() string {

	package2.SomeGlobalfunction()
	return "hello world"

}
