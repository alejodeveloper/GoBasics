package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Race conditions
func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GoRutines\t", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// Set up the race condition
	for i := 0; i <= gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("GoRutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRutines\t", runtime.NumGoroutine())
	fmt.Println("My counter\t", counter)
}
