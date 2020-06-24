package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Printf("-----EXERCISE-----\t%v\n", 1)
	exerciseOne()
	fmt.Printf("-----EXERCISE-----\t%v\n", 2)
	exerciseTwo()
	fmt.Printf("-----EXERCISE-----\t%v\n", 3)
	exerciseThreeAndFour()
	fmt.Printf("-----EXERCISE-----\t%v\n", 5)
	exerciseFive()
}

func exerciseOne() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Something out 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Something out 2")
		wg.Done()
	}()
	wg.Wait()
}

type Person struct {
	name string
	age  int
}

func (p *Person) speak() {
	fmt.Printf("Name: %v and age %v\n", p.name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func exerciseTwo() {
	p1 := Person{"Alexxo", 28}
	saySomething(&p1)
}

func exerciseThreeAndFour() {
	const top = 100
	var wg sync.WaitGroup
	var m sync.Mutex
	counter := 0
	wg.Add(top)
	for i := 0; i < top; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			m.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRutines\t", runtime.NumGoroutine())
	fmt.Println("My counter\t", counter)
}

func exerciseFive() {
	const top = 100
	var wg sync.WaitGroup
	var counter int64
	wg.Add(top)
	for i := 0; i < top; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("GoRutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRutines\t", runtime.NumGoroutine())
	fmt.Println("My counter\t", counter)
}
