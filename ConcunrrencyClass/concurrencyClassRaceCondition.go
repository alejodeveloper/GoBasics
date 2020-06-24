package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Race conditions
func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GoRutines\t", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// Set up the race condition
	for i := 0; i <= gs; i++ {
		go func() {
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("GoRutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRutines\t", runtime.NumGoroutine())
	fmt.Println("My counter\t", counter)
}
