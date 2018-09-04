package main

import "fmt"

func main() {
	// rune = int32
	var x rune = 1<<31 - 1
	fmt.Println(x)
	// 64位系统中, int 为64位的
	var y int = 1<<63 - 1
	fmt.Println(y)
}
