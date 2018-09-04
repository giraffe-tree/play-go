package main

import "fmt"

func main() {

	a, b, c := 1, 2, 3
	fmt.Println(a, b, c)
	fmt.Println(gcd(80, 60))

}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y

	}
	return x
}
