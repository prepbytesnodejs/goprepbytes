package main

import (
	"fmt"
)

type Vacation interface {
	gotoGoa()      //Kartik
	gotoGujrat()   //Amit
	gotoBanglore() // Kavya

}

type Amit struct {
	Packing string
}

func (y2 *Amit) gotoGoa() {

	fmt.Println("Go to goa by train")

}

func (y2 *Amit) gotoGujrat() {

	fmt.Println("Go to Gujrat by train")

}

func (y2 *Amit) gotoBanglore() {

	fmt.Println("Go to Banglore by train")

}

type Karthik struct {
}

func (y2 *Karthik) gotoGoa() {

	fmt.Println("Go to goa by flight")

}

func (y2 *Karthik) gotoGujrat() {

	fmt.Println("Go to Gujrat by flight")

}

func (y2 *Karthik) gotoBanglore() {

	fmt.Println("Go to Banglore by flight")

}

type Kavya struct {
}

func (y2 *Kavya) gotoGoa() {

	fmt.Println("Go to goa by flight")

}

func (y2 *Kavya) gotoGujrat() {

	fmt.Println("Go to Gujrat by train")

}

func (y2 *Kavya) gotoBanglore() {

	fmt.Println("Go to Banglore by Bus")

}
