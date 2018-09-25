package main

import "fmt"

func main() {
	var x uint16 = 1<<16 - 1

	var y = uint8(x)
	fmt.Println(x)
	fmt.Println(y)

}
