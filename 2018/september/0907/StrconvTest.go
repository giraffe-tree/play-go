package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	test1()
}

func test1() {

	s := strconv.FormatBool(true)
	s1, _ := strconv.ParseFloat("1.2323", 32)
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(reflect.TypeOf(s1))

}
