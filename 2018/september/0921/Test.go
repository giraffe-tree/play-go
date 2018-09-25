package main

import "fmt"

func init() {
	fmt.Println("hello init..")
}

func main() {
	fmt.Println(Sum(1, 2))
	fmt.Println(" world")
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
}

//只有当某个函数需要被外部包调用的时候才使用大写字母开头，并遵循 Pascal 命名法；
// 否则就遵循骆驼命名法，即第一个单词的首字母小写，其余单词的首字母大写。
func Sum(a, b int) (x int) {
	return a + b
}
