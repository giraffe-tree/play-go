package main

import "fmt"

func main() {

	var ages = make(map[string]int)
	ages["alice"] = 32

	fmt.Println(ages["alice"])
	var names = map[string]int{
		"chen1": 1,
		"chen2": 2,
		"chen3": 3,
		"chen4": 4,
		"chen5": 5,
		"chen6": 6,
	}
	fmt.Println(names["chen"])

	for _, name := range names {
		// 遍历的顺序不固定
		fmt.Println(name)
	}

	fmt.Println(names["c"])
	names["c"] = 1

	i, ok := names["c"]
	if !ok {
		fmt.Println("不在map i 中")
	} else {
		fmt.Println(i)
	}
	delete(names, "c")
	i, ok = names["c"]
	if !ok {
		fmt.Println("不在map i 中")
	} else {
		fmt.Println(i)
	}
}
