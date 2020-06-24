package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Printf("-----EXERCISE-----\t%v\n", 1)
	exerciseOne()
	fmt.Printf("-----EXERCISE-----\t%v\n", 2)
	exerciseTwo()
	fmt.Printf("-----EXERCISE-----\t%v\n", 3)
	exerciseThree()
	fmt.Printf("-----EXERCISE-----\t%v\n", 4)
	exerciseFour()
}

type Person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func exerciseOne() {
	p1 := Person{"James", "Bond", 32, []string{"Some", "Strings"}}
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalf("There was an error, %v, marshaling the object\n", err)
	}
	fmt.Println(string(bs))
}

func exerciseTwo() {
	p1 := Person{"James", "Bond", 32, []string{"Some", "Strings"}}
	bs, err := toJson(p1)
	if err != nil {
		log.Fatalf("There was an error, %v, marshaling the object\n", err)
	}
	fmt.Println(string(bs))
}

func toJson(p Person) ([]byte, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		err = fmt.Errorf("There was an error, %v, marshaling the object\n", err)
		return []byte{}, err
	}
	return bs, nil
}

type Spy struct {
	name string
}

func (s Spy) Error() string {
	return fmt.Sprintf("The spy %v got discovered\n", s.name)
}

func exerciseThree() {
	james := Spy{"James Bond"}
	foo(james)
	fmt.Println(james)
}

func foo(e error) {
	log.Println(e)
}

func exerciseFour() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

type sqrtError struct {
	lat string
	long string
	err error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Math error:\nLat:\t%v\nLong:\t%v\nError:\t%v\n", se.lat, se.long, se.err)
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		err := errors.New(fmt.Sprintf("Math err, negative number %v\n", n))
		sqrtErr := sqrtError{
			lat:  "-10",
			long: "23",
			err:  err,
		}
		return 0, sqrtErr
	}
	return 41, nil
}