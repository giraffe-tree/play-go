package main

import "fmt"

/**
用于验证 golang 中的值传递

参考: http://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html

*/
func main() {
	var value = 1
	fmt.Printf("主函数中的 value 地址: %p\n", &value)
	// value 值不变, 地址也不变
	//basicType(value)

	// value 值改变, 地址不变
	basicTypeByPointer(&value)
	fmt.Printf("主函数调用 basicType 后 value的值: %d\n", value)
	fmt.Printf("主函数调用 basicType 后 value 的地址: %p\n", &value)

}

func basicType(value int) {
	fmt.Printf("方法中 value 的地址: %p\n", &value)
	value = 8
	fmt.Printf("value 改变后的地址: %p\n", &value)
}

func basicTypeByPointer(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	fmt.Printf("函数里接收到的指针的存储的值 - 即传参的地址是：%p\n", ip)
	*ip = 10
}
