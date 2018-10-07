package main

import (
	"fmt"
	"reflect"
)

func main() {

	x := 1
	x2 := reflect.ValueOf(&x)
	fmt.Println(x2.CanSet())
	x2 = x2.Elem()
	fmt.Println(x2.CanSet())
	x2.SetInt(1123)
	fmt.Println(x2)
	b := x2.Kind() == reflect.Int
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(x))
}
