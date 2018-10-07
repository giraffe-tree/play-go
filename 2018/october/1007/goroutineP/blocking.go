package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

/**
all goroutines are asleep - deadlock!
*/
func main() {
	out := make(chan int)
	//out <- 2
	//go f1(out)
	out <- 2 // 阻塞main goroutine
}
