package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(eve, odd, quit)
	receive(eve, odd, quit)
	fmt.Println("Exiting")
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			odd <- i
		} else {
			even <- i
		}
	}
	close(even)
	close(odd)
	quit <- 0
}
func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("Even channel", v)
		case v := <-odd:
			fmt.Println("Odd channel", v)
		case v := <-quit:
			fmt.Println("Quit channel", v)
			return
		}
	}
}
