package main

import "fmt"

type eatable interface {
	eat()
}

type fruit interface {
	eatable
	water()
}

type Watermelon struct {
	Size int
}

func (w *Watermelon) eat() {
	fmt.Println("eat ... watermelon")
}

func (w *Watermelon) water() {
	fmt.Println("water ... watermelon")
}

func test1(fruit2 fruit) {
	fruit2.water()
}

func test2(fruit2 eatable) {
	fruit2.eat()
}

/**
类型断言可检查值的动态类型, 然而动态类型是什么呢?
 */
func main() {
	watermelon := Watermelon{1}
	var x1 eatable
	test1(&watermelon)
	x1 = &watermelon
	if t, ok := x1.(*Watermelon); ok {
		fmt.Printf("%#v\n", t)
		t.eat()
	}else {
		fmt.Println(t)
	}

	var e eatable = &watermelon
	test2(&watermelon)
	if t, ok := e.(*Watermelon); ok {
		fmt.Printf("%#v", t)
	}
	getType(&watermelon)

}
func getType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case fruit:
		fmt.Println("fruit", v)
	case eatable:
		fmt.Println("eatable", v)
	default:
		fmt.Println("none")
	}
}
