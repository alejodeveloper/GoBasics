package main

import "fmt"

var x int
var y string
var z bool

type alexxo string

var myVar alexxo
var anotherVar string

func main() {
	exerciseOne()
	exerciseTwo()
	exerciseThree()
	exerciseFour()
	exerciseFive()
}


func exerciseOne() {
	a := 42
	b := "James Bond"
	c := true

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// Multiple print in one line
	fmt.Println(a, b, c)
}

func exerciseTwo() {
	fmt.Println(x, y, z)
}

func exerciseThree() {
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%v\t%v\t%v\t ", x, y, z)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}

func exerciseFour() {

	fmt.Println(myVar)
	fmt.Printf("%T\n", myVar)
	myVar = "Sup bitches"
	fmt.Println("---------")
	fmt.Println(myVar)
}


func exerciseFive() {
	fmt.Println(myVar)
	fmt.Printf("%T\n", myVar)
	myVar = "Sup again bitches"
	fmt.Println("---------")
	fmt.Println(myVar)
	anotherVar = string(myVar)

	fmt.Println(anotherVar)
	fmt.Printf("%T\n", anotherVar)

}
