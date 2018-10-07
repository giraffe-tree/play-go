package main

import "fmt"

/**
可以粗略地将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为。
Go 语言中的继承是通过内嵌或组合来实现的，所以可以说，在 Go 语言中，相比较于继承，组合更受青睐。
 */
type innerS struct {
	in1 int
	in2 int
	float64
	b int
}

type outerS struct {
	b    int
	c    float32
	int  // anonymous field
	innerS //anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.innerS.b = 100
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	outer.float64 = 1.0
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	fmt.Printf("outer.float64 is: %f\n", outer.float64)
	fmt.Printf("outer.innerS.b is: %d\n", outer.innerS.b)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10,1.0,2}}
	fmt.Println("outer2 is:", outer2)
}
