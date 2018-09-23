package main

import (
	"fmt"
	"reflect"
)

func main() {

	months := [...]string{1: "January", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May",
		6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	fmt.Println(reflect.TypeOf(months))
	m := months[2:7]

	fmt.Println(reflect.TypeOf(m))
	fmt.Println(m)
	fmt.Println(len(m))

	fmt.Println(cap(m))
	m2 := append(m, "ds")
	fmt.Println(m2)
	fmt.Println(m)
	x1 := copy(months[0:], months[3:])
	fmt.Println(x1)
	fmt.Println(months)
	test()

}
func test() {

}
