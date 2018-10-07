package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type Point2 struct {
	x, y float64
}

func (p *Point2) Abs() float64 {
	return math.Sqrt(100)
}

/**
因为一个结构体可以嵌入多个匿名类型，所以实际上我们可以有一个简单版本的多重继承，
就像：type Child struct { Father; Mother}。在第 10.6.7 节中会进一步讨论这个问题。
 */
//func (p *NamedPoint) Abs() float64 {
//	return math.Sqrt(p.point.x*p.point.x + p.point.y*p.point.y +400)
//}

type NamedPoint struct {
	Point
	point2 Point2
	name string
}

/**
当多个同名函数通过内嵌类型出现在同一个结构体时, 会触发编译报错 ambiguous selector type.method
 */
func main() {
	n := &NamedPoint{Point{3, 4}, Point2{1, 2}, "hello"}
	fmt.Println(n.Abs())     // 打印5
	fmt.Println(n.point2.Abs()) // 打印5
	fmt.Println(n.Point.Abs()) // 打印5
}
