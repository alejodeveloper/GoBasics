package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type User struct {
	First string
	Age   int
}

type Spy struct {
	First  string   `json:"First"`
	Last   string   `json:"Last"`
	Age    int      `json:"Age"`
	Saying []string `json:"Saying"`
}

type Person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

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
}

func exerciseOne() {
	u1 := User{"James", 32}
	u2 := User{"Moneypenny", 25}
	u3 := User{"M", 54}
	users := []User{u1, u2, u3}
	fmt.Println(users)
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonUsers))
}

func exerciseTwo() {
	jsonSpies := `[{"First":"James","Last": "Bond","Age":32, "Saying": ["Some", "Various", "Strings"]},{"First":"Miss","Last":"Moneypenny","Age":25, "Saying": ["Some", "Various", "Anothers","Strings"]},{"First":"M","Last":"Who knows?","Age":54, "Saying": ["Some", "Various", "Anothers","Strings"]}]`
	jsonResponse := []byte(jsonSpies)
	var spies []Spy
	err := json.Unmarshal(jsonResponse, &spies)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(spies)
}

func exerciseThree() {
	p1 := Person{"James", "Bond", 32, []string{"Some", "Strings"}}
	p2 := Person{"Miss", "Moneypenny", 25, []string{"More", "Strings"}}
	p3 := Person{"M", "Hhhhh", 25, []string{"Wait", "There is", "More?"}}
	people := []Person{p1, p2, p3}
	fmt.Println(people)
	err := json.NewEncoder(os.Stdout).Encode(people)
	if err != nil {
		fmt.Println(err)
	}
}

func exerciseFour() {
	myNumbers := []int{55, 89, 55, 63, 74, 11, 258, 6, 12}
	myStrings := []string{"hello", "Kasdas", "Tyu", "MM", "asd", "ADF", "Pedro", "Xiaomi", "China", "Honda", "Kawasaki", "Ducati"}
	fmt.Println(myNumbers)
	fmt.Println(myStrings)
	sort.Ints(myNumbers)
	sort.Strings(myStrings)
	fmt.Println(myNumbers)
	fmt.Println(myStrings)
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByLast []Person

func (a ByLast) Len() int {
	return len(a)
}
func (a ByLast) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByLast) Less(i, j int) bool {
	return a[i].Last < a[j].Last
}

func exerciseFive() {
	p1 := Person{"James", "Bond", 32, []string{"Hey", "Martini", "Kill", "Balls", "Strings"}}
	p2 := Person{"Miss", "Moneypenny", 25, []string{"More", "Strings", "Bond", "Mr", "Cosmopolitan"}}
	p3 := Person{"M", "Hhhhh", 25, []string{"Wait", "There is", "More?", "UK", "The Queen", "Your majesty"}}
	people := []Person{p1, p2, p3}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByLast(people))
	fmt.Println(people)
	for _, value := range people {
		sort.Strings(value.Sayings)
	}
	fmt.Println(people)
}
