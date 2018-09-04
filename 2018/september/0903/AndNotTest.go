package main

import "fmt"

func main() {

	// 1111
	var x = 0x0
	// 1010
	var y = 0x1

	// 将运算符左边数据相异的位保留，相同位清零。

	fmt.Println(x &^ y) // 0 &^ 1 = 0
	fmt.Println(x &^ x) // 0 &^ 0 = 0
	fmt.Println(y &^ y) // 1 &^ 1 = 0
	fmt.Println(y &^ x) // 1 &^ 0 = 0

}
