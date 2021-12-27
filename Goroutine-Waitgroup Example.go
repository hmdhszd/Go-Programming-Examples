package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin CPUs:", runtime.NumCPU())
	fmt.Println("begin Goroutines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello from Up")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from Down")
		wg.Done()
	}()

	fmt.Println("Middle CPUs:", runtime.NumCPU())
	fmt.Println("Middle Goroutines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")

	fmt.Println("End CPUs:", runtime.NumCPU())
	fmt.Println("End Goroutines", runtime.NumGoroutine())
}
