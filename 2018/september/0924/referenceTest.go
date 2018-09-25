package main

import "fmt"

/**
用于测试 golang 中的引用问题
由于在 java 中都是值传递, 想看看在 go 中是怎么样的
*/
func main() {
	apple := &Apple{"cc", 1}
	fmt.Printf("原始 apple 地址:%p", &apple)
	fmt.Println(apple.name, apple.size) // cc 1

	fmt.Printf("原始 apple 地址:%p\n", &apple)
	test1(*apple)
	fmt.Println(apple.name, apple.size) // cc 1

	fmt.Printf("原始 apple 地址:%p\n", &apple)
	test2(apple)
	fmt.Println(apple.name, apple.size) // chen 10

	fmt.Printf("原始 apple 地址:%p\n", &apple)
	apple = &Apple{"cc", 1}
	test3(apple)
	fmt.Println(apple.name, apple.size) // chen 333

	fmt.Printf("\n原始 apple 地址:%p\n", &apple)
	apple = &Apple{"cc", 1}
	test4(apple)
	fmt.Println(apple.name, apple.size) // cc 1
	fmt.Printf("原始 apple 地址:%p\n", &apple)

	ints := make(map[string]int)
	fmt.Println(ints)
}

// 引用了变量 apple 的一个副本, 及时在函数内部改变对象, 也不能影响原对象, 故原对象状态不改变
func test1(apple Apple) {
	fmt.Printf("test1 apple 地址:%p", &apple)
	apple.name = "chen"
	apple.size = 10
}

// 将变量 apple 的指针传递给函数, 隐式地将 apple 的地址也传递给了函数
// 所以函数执行后, 原变量的状态发生了改变
func test2(apple *Apple) {
	fmt.Printf("test2 apple 地址:%p", &apple)
	apple.name = "chen"
	apple.size = 10
}

// 对变量进行重新赋值
func test3(apple *Apple) {
	fmt.Printf("test3 apple 地址:%p\n", &apple)
	*apple = Apple{"chen", 333}
	fmt.Printf(" apple 地址:%p\n", &apple)
}

func test4(apple *Apple) {
	fmt.Printf("test4 方法中 apple 值:%p\n", apple)
	fmt.Printf(" apple 地址:%p\n", &apple)
	apple = &Apple{"xxx", 444}
	fmt.Printf(" apple 地址:%p\n", &apple)
	fmt.Printf(" apple 值:%p\n", apple)
}

type Apple struct {
	name string
	size int
}
