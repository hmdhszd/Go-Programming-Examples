package main

import "fmt"

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)

	fmt.Println("About to exit...")
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(quit)

}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("From the even channel:", v)
		case v := <-odd:
			fmt.Println("From the odd channel:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("!ok --> i, ok =", i, ok)
				return
			} else {
				fmt.Println("ok --> i =", i)
			}

		}
	}

}
