package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
	int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("\nYou can make it!")
	} else {
		fmt.Println("\nNeed to fuel up!")
	}
}
func main() {
	var myEngine electricEngine = electricEngine{25, 15}
	// myEngine.mpg = 25        // Can rename the field
	var myEngine2 = struct { // Anonymous Struct, not reusable
		mpg     uint8
		gallons uint8
	}{25, 15}
	var myEngine3 = struct { // have to rename again if I want to reuse it
		mpg     uint8
		gallons uint8
	}{25, 15}
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name, myEngine.int)
	fmt.Println(myEngine2.mpg, myEngine2.gallons)
	fmt.Println(myEngine3.mpg, myEngine3.gallons)
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())
	canMakeIt(myEngine, 50)
}
