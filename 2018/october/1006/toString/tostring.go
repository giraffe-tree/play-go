package main

import (
	"fmt"
	"runtime"
)

type TwoInts struct {
	a int
	b int
}

/**

当你广泛使用一个自定义类型时，最好为它定义 String()方法。
从上面的例子也可以看到，格式化描述符 %T 会给出类型的完全规格，%#v 会给出实例的完整输出，
包括它的字段（在程序自动生成 Go 代码时也很有用）。
 */
func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	// 值表示
	// 类型表示
	// 类型+值
	// 字段+值

	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
	fmt.Printf("two1 is: %+v\n", two1)
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
 	fmt.Printf("%v kb\n", stats.Alloc/1024)
}

// 这里定义的是 值的表示,所有 %v 的表示都会改变
//func (tn *TwoInts) String() string {
//	return "(" + strconv.Itoa(tn.a) + " - " + strconv.Itoa(tn.b) + ")"
//}
