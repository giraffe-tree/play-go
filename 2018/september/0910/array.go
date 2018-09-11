package main

import "fmt"

func main() {

	a := [...]int{1,2}
	b :=[2]int{2,3}
	c :=[2]int{1,2}
	//d :=[3]int{1,2}

	fmt.Println(a==b)
	fmt.Println(a==c)
	fmt.Println(b==c)
	//fmt.Println(a==d)

}
