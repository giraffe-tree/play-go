package main

import "fmt"

func main() {
	var x uint8 = 255
	y := 2
	var z = 1<<63 - 1
	y, z = z, y
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(&z)

}
