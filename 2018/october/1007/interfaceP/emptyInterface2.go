package interfaceP

import "fmt"

func main() {
	i2 := make([]interface{}, 10)
	i2[0] = "string"
	i2[1] = 1

	fmt.Println(i2)
	paper := Paper{1}
	f3(&paper)
}

type Ptype interface {
	print()
}

type Wtype interface {
	write()
}

func f3(x Xempty) {
	if t,ok:=x.(Ptype);ok{
		t.print()
	}
	if t,ok :=x.(Wtype);ok{
		t.write()
	}
}

type Xempty interface {
}

type Paper struct {
	size int
}

func (p Paper) print() {
	fmt.Println("paper print ...")
}
