package main

import "fmt"

func main() {
	var x float32 = 1 << 24

	// 由于 float32 会导致精度缺失, 下面的表达式输出 true
	fmt.Println(x == x+1)
}
