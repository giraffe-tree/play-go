package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	square := Square{122}
	//var x Shaper
	//x = square
	//fmt.Printf("The square has area: %f\n", square.Area())

	//square := new(Square)
	//square.side = 5

	var areaIntf Shaper
	areaIntf = &square
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

}
