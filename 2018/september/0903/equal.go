package main

import "fmt"

func main() {

	var x = 1
	var y = 1
	// 数值相同
	fmt.Println(x == y)
	// 但地址不同
	fmt.Println(&x == &y)

}
