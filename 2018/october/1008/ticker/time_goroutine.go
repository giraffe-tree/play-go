package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e9)
	boom := time.After(3e9)
	for {
		select {
		case v := <-tick:
			fmt.Println("tick.", v)
		case v := <-boom:
			fmt.Println("BOOM!", v)
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e8)
		}
	}
}
