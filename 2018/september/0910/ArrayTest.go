package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
)

func main() {

	var a [3]int = [3]int{2, 3}
	fmt.Println(a[2])
	i := [3]int{1, 2, 3}
	fmt.Printf("%T\n", i)
	symbol := [10]string{USD: "ds", 2: "dsa"}
	fmt.Println(symbol[0])
	fmt.Println(symbol[USD])

}
