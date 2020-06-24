package main

import (
	"fmt"
)

func main() {
	fmt.Printf("----- Exercise ----------\t %v \n", 1)
	exerciseOne()
	fmt.Printf("----- Exercise ----------\t %v \n", 2)
	exerciseTwo()

}

func exerciseOne() {
	x := 56
	fmt.Printf("Some value %v has the pointer %v/\n", x, &x)
}

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {

	(*p).first = "We change your name Broh"
	(*p).last = "To this"
	(*p).age = 40
}

func exerciseTwo() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Printf("The person before %v\n", p)
	changeMe(&p)
	fmt.Printf("The person after %v\n", p)

}
