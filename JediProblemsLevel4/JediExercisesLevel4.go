package main

import (
	"fmt"
)

func main() {
	fmt.Printf("----------\nExercice %v\n", 2)
	exerciseTwo()
	fmt.Printf("----------\nExercice %v\n", 3)
	exerciseThree()
	fmt.Printf("----------\nExercice %v\n", 4)
	exerciseFour()
	fmt.Printf("----------\nExercice %v\n", 5)
	exerciseFive()
	fmt.Printf("----------\nExercice %v\n", 6)
	exerciseSix()
	fmt.Printf("----------\nExercice %v\n", 7)
	exerciseSeven()
	fmt.Printf("----------\nExercice %v\n", 8)
	exerciseEight()
	fmt.Printf("----------\nExercice %v\n", 9)
	exerciseNine()
	fmt.Printf("----------\nExercice %v\n", 10)
	exerciseTen()
}
func exerciseTwo() {
	x := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i, value := range x {
		fmt.Printf("Over the index %v we have the value %v\n", i, value)
	}
	fmt.Printf("%T\n", x)
}

func exerciseThree() {
	x := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Printf("%v\n", x[:5])
	fmt.Printf("%v\n", x[4:])
	fmt.Printf("%v\n", x[4:9])
	fmt.Printf("%v\n", x[2:7])
}

func exerciseFour() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	fmt.Println(x)
	x = append(x, y...)
	fmt.Println(x)
}

func exerciseFive() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func exerciseSix() {
	states := []string{"Alabama", "Alaska", "Arizona", "Arkansas",
		"California",
		"Colorado",
		"Connecticut",
		"Delaware",
		"Florida",
		"Georgia",
		"Hawaii",
		"Idaho",
		"IllinoisIndiana",
		"Iowa",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Maine",
		"Maryland",
		"Massachusetts",
		"Michigan",
		"Minnesota",
		"Mississippi",
		"Missouri",
		"MontanaNebraska",
		"Nevada",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York",
		"North Carolina",
		"North Dakota",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"PennsylvaniaRhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Vermont",
		"Virginia",
		"Washington",
		"West Virginia",
		"Wisconsin",
		"Wyoming"}
	fmt.Printf("The capacity of the slice is %v\n", cap(states))
	fmt.Printf("The len of the slice is %v\n", len(states))
	for i := 0; i < len(states); i++ {
		fmt.Printf("State %v is in the index %v\n", states[i], i)
	}
}

func exerciseSeven() {
	x := [][]string { {"James","Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Hellooooo, Mr Bond"}}
	for rowIndex, value := range x {
		for columnIndex, data := range value {
			fmt.Printf("The index row %v and the column %v has the value %v\n", rowIndex, columnIndex, data)
		}
	}
}

func exerciseEight() {
	x := map[string] []string{
		"bond_james": {"Kill", "Ice cream", "Martinis"},
		"moneypenny_miss": {"Ice cream", "Mr Bond"},
		"no_dr": {"Computer Science", "Kill"},
	}
	fmt.Println(x)
	fmt.Println("*/*/*/*")
	for key, value := range x {
		for _, data := range value {
			fmt.Printf("First name: %v likes: %v\n", key, data)
		}

	}
}

func exerciseNine() {
	x := map[string] []string{
		"bond_james": {"Kill", "Ice cream", "Martinis"},
		"moneypenny_miss": {"Ice cream", "Mr Bond"},
		"no_dr": {"Computer Science", "Kill"},
	}
	fmt.Println(x)
	fmt.Println("*/*/*/*")
	for key, value := range x {
		for _, data := range value {
			fmt.Printf("First name: %v likes: %v\n", key, data)
		}

	}
	x["miss_m"] = []string{"Building gadgets", "Mocking of bond", "Sunsets"}
	fmt.Println("*/*/*/*")
	fmt.Println("With miss M")
	for key, value := range x {
		for _, data := range value {
			fmt.Printf("First name: %v likes: %v\n", key, data)
		}

	}

}
func exerciseTen() {
	x := map[string] []string{
		"bond_james": {"Kill", "Ice cream", "Martinis"},
		"moneypenny_miss": {"Ice cream", "Mr Bond"},
		"no_dr": {"Computer Science", "Kill"},
	}
	fmt.Println(x)
	fmt.Println("*/*/*/*")
	for key, value := range x {
		for _, data := range value {
			fmt.Printf("First name: %v likes: %v\n", key, data)
		}

	}
	delete(x, "no_dr")
	fmt.Println("*/*/*/*")
	fmt.Println("Without MR No")
	for key, value := range x {
		for _, data := range value {
			fmt.Printf("First name: %v likes: %v\n", key, data)
		}

	}

}