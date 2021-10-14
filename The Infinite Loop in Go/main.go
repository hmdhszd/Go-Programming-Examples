package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"myapp/mylogger"
)

func main() {

	//read input from the user 5 times and write it to a log

	reader := bufio.NewReader(os.Stdin)

	ch := make(chan string)

	go mylogger.ListenForLog(ch)

	fmt.Println("Enter something...")

	for i := 0; i < 5; i++ {
		fmt.Println("-->")
		input, _ := reader.ReadString('\n')

		ch <- input
		time.Sleep(time.Second)
	}

}
