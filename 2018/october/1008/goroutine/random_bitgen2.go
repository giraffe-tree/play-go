package goroutine

import "fmt"

/**
简单的来说使用 select 可以发送/接受chan中的数据
 */
func main() {
	c := make(chan int)
	// consumer:
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	// producer:
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}

}
