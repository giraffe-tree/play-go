package main

import "fmt"

func main() {
	//for x := 1; x < 10; x++ {
	//	fmt.Println(f() == f(),*f())
	//}
	var v = 1
	incr(&v)
	fmt.Println(v)
}

var p = f()

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
