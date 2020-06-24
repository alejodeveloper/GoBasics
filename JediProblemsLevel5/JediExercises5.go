package main

import "fmt"

func main() {
	fmt.Printf("-------------\nExercise %v\n", 1)
	exerciseOne()
	fmt.Printf("-------------\nExercise %v\n", 2)
	exerciseTwo()
	fmt.Printf("-------------\nExercise %v\n", 3)
	exerciseThree()
	fmt.Printf("-------------\nExercise %v\n", 4)
	exerciseFour()
}

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func exerciseOne() {
	personA := person{
		firstName: "Some Name",
		lastName:  "Some Last Name",
		iceCream:  []string{"Chocolate", "Strawberry"},
	}
	personB := person{
		firstName: "Some Another Name",
		lastName:  "Some Another Last Name",
		iceCream:  []string{"Brownie", "Vanilla"},
	}

	fmt.Println(personA.firstName)
	fmt.Println(personA.lastName)
	for _, iceCream := range personA.iceCream {
		fmt.Printf("Ice Cream Flavor %v\n", iceCream)
	}
	fmt.Println("********************")
	fmt.Println("--------------------")
	fmt.Println("********************")
	fmt.Println(personB.firstName)
	fmt.Println(personB.lastName)
	for _, iceCream := range personB.iceCream {
		fmt.Printf("Ice Cream Flavor %v\n", iceCream)
	}
}

func exerciseTwo() {
	personA := person{
		firstName: "Some Name",
		lastName:  "Some Last Name",
		iceCream:  []string{"Chocolate", "Strawberry"},
	}
	personB := person{
		firstName: "Some Another Name",
		lastName:  "Some Another Last Name",
		iceCream:  []string{"Brownie", "Vanilla"},
	}
	x := map[string]person{
		personA.lastName: personA,
		personB.lastName: personB,
	}

	for key, value := range x {
		fmt.Printf("The %v belongs to %v \n", key, value)
	}
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func exerciseThree() {
	mySedan := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		luxury: true,
	}
	myTruct := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	fmt.Println(mySedan)
	fmt.Println(myTruct)
}

func exerciseFour() {
	myMap := map[string]string{
		"SomeKey": "borracho",
		"Dice":    "Say",
	}
	mySlice := []string{"a", "b", "c"}
	x := struct {
		m       map[string]string
		mySlice []string
	}{
		m:       myMap,
		mySlice: mySlice,
	}

	fmt.Println(x)

}

// callback when a function is pass as parameter in another function
func sumOdd(f func(... int) int, y ... int) int {
	var x []int
	for _, number := range y {
		if number % 2 != 0 {
			x = append(x, number)
		}
	}
	return f(x...)
}
// enclosure
func incrementor() func() int {
	// we are enclosing the scope of x to a chunk of code
	var x int
	return func() int {
		x++
		return x
	}
}

func factorialRecursive(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorialRecursive(x-1)
}

func factorialIterative(x int) int {
	factorial := 1
	for i:=x; i>=1; i -- {
		factorial = factorial * i
	}
	return factorial
}