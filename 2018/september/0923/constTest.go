package main

import "fmt"

const Pi = 3.14159265358979

// 反斜杠好像不不能使用 176568075500134360255254120680009

const (
	Sunday    = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {

	fmt.Println(Pi)
	fmt.Println(Sunday)
	fmt.Println(Thursday)
	fmt.Println(Monday)
}
