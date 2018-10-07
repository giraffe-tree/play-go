package main

import (
	"fmt"
	"time"
)

func main() {
	sendChan := make(chan int)
	receiveChan := make(chan<- string)
	go processChannel(sendChan, receiveChan)

	go func() {
		for i := 0; ; i++ {
			sendChan <- i
		}
	}()
	time.Sleep(2e9)
}

func processChannel(in <-chan int, out chan<- string) {
	for inValue := range in {
		result := string(inValue) /// processing inValue
		fmt.Println(result)
		out <- result
	}
}
