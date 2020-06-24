package main

import (
	"fmt"
	"sort"
)

func main() {
	sortExample()
}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.name, p.age)
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}
func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].name < a[j].name
}

func sortExample() {
	p1 := Person{"James", 32}
	p2 := Person{"Q", 65}
	p3 := Person{"M", 54}
	p4 := Person{"Moneypenny", 25}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
