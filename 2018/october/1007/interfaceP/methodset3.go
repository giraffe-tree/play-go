package interfaceP

import "fmt"

/**
如果你指定了 函数要接受一个指针, 那么它只能接受指针,
而如果你指定函数接受一个值, 那么它既可以接受值, 又可以接受指针,
因为指针会被自动解引用
 */

func main() {
	man := new(Man)
	//man.money = 1000
	//man.height = 186
	//man.weight = 66
	readBook(man)
	man.read()
	//man := Man{186, 66, 1000}
	//man.read()
	//readBook(&man)

	var m Man
	fmt.Println(m.weight)
	readBook(&m)
}

type read interface {
	read()
}

type buyBook interface {
	buyBook(b *Book)
}

type Man struct {
	height float64
	weight float64
	money  int
}

type Book struct {
	name  string
	size  int
	money int
}

func (m *Man) read() {
	fmt.Println("read with 值 man")
}

func readBook(r read) {
	r.read()
}
