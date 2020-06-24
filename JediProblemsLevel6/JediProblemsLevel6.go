package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("----- Exercise ----------\t %v \n", 1)
	exerciseOne()
	fmt.Printf("----- Exercise ----------\t %v \n", 2)
	exerciseTwo()
	fmt.Printf("----- Exercise ----------\t %v \n", 3)
	exerciseThree()
	fmt.Printf("----- Exercise ----------\t %v \n", 4)
	exerciseFour()
	fmt.Printf("----- Exercise ----------\t %v \n", 5)
	exerciseFive()
	fmt.Printf("----- Exercise ----------\t %v \n", 6)
	exerciseSix()
	fmt.Printf("----- Exercise ----------\t %v \n", 7)
	exerciseSeven()
	fmt.Printf("----- Exercise ----------\t %v \n", 8)
	exerciseEight()
	fmt.Printf("----- Exercise ----------\t %v \n", 9)
	exerciseNine()
	fmt.Printf("----- Exercise ----------\t %v \n", 10)
	exerciseTen()
}

func exerciseOne() {
	x := fooOne()
	y, s := barOne()
	fmt.Println("Foo returns the int")
	fmt.Println(x)
	fmt.Println("Bar returns the int")
	fmt.Println(y)
	fmt.Println("And the string")
	fmt.Println(s)

}

func fooOne() int {
	return 42
}

func barOne() (int, string) {
	return 40, "Some random string"
}

func exerciseTwo() {
	myNumbers := []int{1, 2, 3, 4, 5}

	x := fooTwo(myNumbers...)
	y := barTwo(myNumbers)
	fmt.Println("Foo returns the sum")
	fmt.Println(x)
	fmt.Println("Bar returns the sum")
	fmt.Println(y)

}

func fooTwo(n ...int) int {
	sum := 0
	for _, value := range n {
		sum += value
	}
	return sum
}

func barTwo(n []int) int {
	sum := 0
	for _, value := range n {
		sum += value
	}
	return sum
}

func exerciseThree() {
	myNumbers := []int{1, 2, 3, 4, 5}
	defer fmt.Println(fooOne())
	defer fmt.Println(fooTwo(myNumbers...))
	fmt.Println(barOne())
	fmt.Println(barTwo(myNumbers))

}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello, my name is %v %v and my age is %d\n", p.first, p.last, p.age)
}

func exerciseFour() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p.speak()
}

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type square struct {
	size float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.size * s.size
}

func info(s shape) {
	fmt.Printf("The square area is %v\n", s.area())
}

func exerciseFive() {
	s := square{
		size: 5,
	}

	c := circle{
		radius: 2,
	}

	info(s)
	info(c)
}

func exerciseSix() {
	x := 2
	func(n int) int {
		return n * 2
	}(x)

	fmt.Printf("After a few secs %v\n", x)

}

func exerciseSeven() {
	x := func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}
	x()
}

func mySum() func(...int) int {
	return func(n ...int) int {
		sum := 0
		for _, value := range n {
			sum += value
		}
		return sum
	}
}

func exerciseEight() {
	sum := mySum()
	integers := []int{1, 2, 3, 4}
	fmt.Printf("Sum my integers %v and the result is %v\n", integers, sum(integers...))
}

func sumOdds(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func anotherSum(f func(...int) int) {
	var n []int
	for len(n) < 100 {
		n = append(n, 1)
	}
	fmt.Println("So we are going to sum into a callback 100 1's")
	fmt.Printf("And the result is %v\n", f(n...))

}

func exerciseNine() {
	anotherSum(sumOdds)
}

func exerciseTen() {
	var x int
	func(n int) {
		x = 2 + n
	}(x)
	fmt.Println(x)
}
