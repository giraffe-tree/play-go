package main

import (
	"fmt"
	"time"
)

/**
channel idiom 简化版
*/
func main() {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch)
		}
		// 以下代码不会运行
		fmt.Println("test")
		for v := range ch {
			fmt.Printf("%v\n", v)
		}
	}()
	time.Sleep(1e9)
}
