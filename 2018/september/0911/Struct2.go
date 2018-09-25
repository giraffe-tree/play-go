package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	var x tree
	x.value = 1
	x.left = nil
	var y tree
	x.right = &y
	y.value = 3
	fmt.Println(x.right.value)
	z := tree{value: 1, left: nil, right: &y}
	fmt.Println(z)

	pp := &tree{111, nil, nil}
	p2 := new(tree)
	*p2 = tree{12, nil, nil}
	fmt.Println(pp)
	fmt.Println(p2)

}
