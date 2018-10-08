package goroutine

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go test(ch1, ch2)
	go test2(ch2, ch1)
	time.Sleep(1e4)
}
func test(ch1 chan int, ch2 chan int) {
	fmt.Println("begin")
	i := 0
	for {
		i2 := <-ch1
		i += i2
		fmt.Println(i2)
		ch2 <- i
	}
}

func test2(ch1 chan int, ch2 chan int) {
	ch2 <- 1
	i := 0
	for {
		i2 := <-ch1
		i += i2
		fmt.Println(i2)
		ch2 <- i
	}
}
