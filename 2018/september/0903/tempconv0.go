package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BiilingC      Celsius = 100
)

func init() {

	fmt.Println("hello world")
	println("hesee")
}

func main() {
	fmt.Println(CToF(100))
	fmt.Printf("dsdsds")
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
