package main

import (
	"fmt"
)

// Interface says what you have to do it doesnt says how you have to do

//decalaration
//definition /implementation

// interface is a type /blueprint that lays out some rules that if any concrete type follows become the part of interface

// Zara ,HRX ,Adidas

type socket interface {
	fourPinpowerPlug() //decalaraing method
	//powerSocket()
}

//concrete types

// you can use them as value int ,float ,struct

type MobileDevice struct {
	Name string
}

func (m MobileDevice) fourPinpowerPlug() {

	fmt.Println("I am implementing 4 pin power interface")

}

func (m MobileDevice) powerSocket() {

	fmt.Println("I am implementing 4 pin power interface")

}

type MixerGrinder struct {
	Name string
}

func (m MixerGrinder) powerSocket() {

	fmt.Println("I am  mixer grinder and implementing 4 pin power interface")

}

func (m MixerGrinder) fourPinpowerPlug() {
	// ---it depends on you what you want to do here

	fmt.Println("I am  mixer grinder and implementing 4 pin power interface")

}

type Laptop struct {
	Brand string
}

func (l Laptop) fourPinpowerPlug() {

	fmt.Println("I am  laptop and implementing 4 pin power interface")

}

func (m Laptop) powerSocket() {

	fmt.Println("I am  mixer grinder and implementing 4 pin power interface")

}

func main() {

	myDevices := []socket{}

	myDevices = append(myDevices, MixerGrinder{Name: "Philips"}, MobileDevice{Name: "Apple"}, Laptop{
		Brand: "Acer",
	})
	fmt.Println(myDevices)

	// fmt.Println(myDevices[0]==myDevices[1])

	//fmt.Println("Hello, World!")
}
