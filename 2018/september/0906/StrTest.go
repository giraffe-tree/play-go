package main

import "fmt"

func main() {

	s := "helloworld"
	fmt.Println(len(s))
	fmt.Println(s[9])
	fmt.Println(s[:])
	// 取左不取右
	fmt.Println(s[1:])
	a := s
	b := "hello, world"
	// go 中也有常量池?
	fmt.Println(a == b)
	fmt.Println(s == b)
	fmt.Println(a == s)
	fmt.Println(&a)

	c := "hello"
	e := "aaa"
	d := "world"

	fmt.Println(&c)
	fmt.Println(&d)
	fmt.Println(&e)

}
