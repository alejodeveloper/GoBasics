package main

import (
	"fmt"
)

func main() {
	exerciseOne()
	exerciseTwo()
	exerciseThree()
	exerciseFour()
	exerciseFive()
	exerciseSix()
	exerciseSeven()
	exerciseEight()
	exerciseNine()
}

func exerciseOne() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exerciseTwo() {
	fmt.Println("-----------------------")
	fmt.Println("EXERCISE #2")
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
func exerciseThree() {
	fmt.Println("-----------------------")
	fmt.Println("EXERCISE #3")
	y := 1992
	for y <= 2020 {
		fmt.Println(y)
		y++
	}
}

func exerciseFour() {
	fmt.Println("-----------------------")
	fmt.Println("EXERCISE #4")
	y := 1992
	for {
		fmt.Println(y)
		y++
		if y == 2020 {
			break
		}
	}
}

func exerciseFive() {
	fmt.Println("-----------------------")
	fmt.Println("EXERCISE #5")

	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

func exerciseSix() {
	fmt.Println("-----------------------")
	fmt.Println("EXERCISE #6")
	x := 100
	y := 52
	if (x + y) != 200 {
		fmt.Println("AHOLE HI")
	}
}
func exerciseSeven() {
	fmt.Println("-----------------------")
	fmt.Println("EXERCISE #7")
	x := 100
	y := 52
	if (x + y) == 200 {
		fmt.Println("JuI")
	} else if x == 52 {
		fmt.Println("Shouldn be here")
	} else {
		fmt.Printf("My name is %v and %v", x, y)
	}
}

func exerciseEight() {
	fmt.Println("-----------------------")
	fmt.Println("EXERCISE #8")
	x := 100
	y := 52
	switch {
	case false:
		println("This is nonsense")

	case x+y == 200:
		fmt.Println("Is 8 not 7")
	default:
		println("Im trying new stuff")
	}
}

func exerciseNine() {
	fmt.Println("-----------------------")
	fmt.Println("EXERCISE #9")
	var favSport interface{} = "SOME STRING"
	switch v := favSport.(type) {
	case string:
		fmt.Println("The var %v a string", v)

	case int:
		fmt.Println("Is 8 not 7")
	default:
		println("Im trying new stuff")
	}
}
