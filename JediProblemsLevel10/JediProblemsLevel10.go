package main

import (
	"fmt"
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
	fmt.Printf("-----EXERCISE-----\t%v\n", 5)
	exerciseFive()
	fmt.Printf("-----EXERCISE-----\t%v\n", 6)
	exerciseSix()
	fmt.Printf("-----EXERCISE-----\t%v\n", 7)
	exerciseSeven()
}

func exerciseOne() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}

func exerciseTwo() {
	cs := make(chan int)
	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func exerciseThree() {
	c := gen()
	receive(c)
	fmt.Printf("About to exit \n")
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("Output from channel", v)
	}
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receiveFour(numbers, quit <-chan int) {
	for {
		select {
		case v := <-numbers:
			fmt.Println("Numbers channel", v)
		case v := <-quit:
			fmt.Println("Out", v)
			return
		}
	}
}

func genFour(quit chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		quit <- 0
		close(c)
	}()

	return c
}

func exerciseFour() {
	q := make(chan int)
	c := genFour(q)
	receiveFour(c, q)
	fmt.Println("About to go")
}

func exerciseFive() {
	c := make(chan int)
	go func() {
		c <- 41
	}()
	v, ok := <-c
	fmt.Printf("Channel\t%v\n OK\t%v\n", v, ok)
	close(c)
	v, ok = <-c
	fmt.Printf("Channel\t%v\n OK\t%v\n", v, ok)

}

func exerciseSix() {
	c := make(chan int)
	go send(c)
	receiver(c)
	fmt.Println("About to exit")
}

func send(numbers chan int) chan int {
	for i := 0; i < 100; i++ {
		numbers <- i
	}
	close(numbers)
	return numbers
}

func receiver(numbers chan int) {
	for number := range numbers {
		fmt.Println(number)
	}
}

func exerciseSeven() {
	pipe := make(chan int)
	for i := 0; i < 10; i++ {
		fmt.Println("Adding")
		go addNumber(pipe)
	}
	getNumbers(pipe)
}

func addNumber(pipe chan<- int) {
	numbers := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	convergeNumbers(numbers, pipe)
}

func convergeNumbers(numbers <-chan int, pipe chan<- int) {
	for number := range numbers {
		fmt.Println("Adding of ", number)
		go func(n int) {
			pipe <- n
		}(number)
	}
}

func getNumbers(numbers <- chan int) {
	go func() {
		for number := range numbers {
			fmt.Println("Getting of ", number)
		}
	}()
	_,ok := <- numbers
	if !ok {
		fmt.Println("END")
	}
}
