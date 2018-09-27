package main

import (
	"fmt"
	"strings"
)

func main() {
	//testReplace("asdas c jdladj ai n nen nqw")
	//test2("abcdef")

	test3()
}

func testReplace(s string) {
	var a = strings.Replace(s, " ", "", -1)
	fmt.Println(a)
}

func test2(s string) {
	fmt.Println(s[1:4])
}

func test3() {

	a := "ab"
	b := "a" + "b"
	fmt.Println(a == b)

}
