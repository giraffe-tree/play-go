package goroutine

import (
	"fmt"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			ch1 <- 1
		}
	}()

	go func() {
		for {
			ch2 <- 0
		}
	}()
	for {
		select {
		case <-ch1:
			fmt.Println(<-ch1)
		case <-ch2:
			fmt.Println(<-ch2)
		}
	}

}
