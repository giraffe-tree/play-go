package main

import "fmt"

func main() {
	//使用 panic 语句（第 13.2 节）来获取栈跟踪信息（直到 panic 时所有被调用函数的列表）。
	//
	//使用关键字 defer 来跟踪代码执行过程（第 6.4 节）。
	fmt.Printf("%T", "chen")
}
