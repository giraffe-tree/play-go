package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	W11 = -2.1
	W12 = 0.23
	W13 = -2.5
	W21 = -2.1
	W22 = -2.6
	W23 = 0.2
	B1  = -0.17
	B2  = -0.14
	B3  = 0.22
	Q1  = 3.0
	Q2  = -2.6
	Q3  = -2.5
)

func main() {
	num1 := 0
	num2 := 0
	strings := make([]string, 100)
	snum := 0
	for i := 0; i < 100; i++ {
		f1 := rand.Float32()*12 - 6
		f2 := rand.Float32()*12 - 6
		f := calc(f1, f2)
		flag := 1
		if f < 0 {
			flag = 0
		}
		if (f1*f2 >= 0 && flag == 1) || (f1*f2 < 0 && flag == 0) {
			num1++
		} else {
			num2++
			s := strconv.FormatFloat(float64(f1), 'f', -1, 32)
			s += "," + strconv.FormatFloat(float64(f2), 'f', -1, 32)
			s += " -> "
			s += strconv.FormatFloat(float64(f), 'f', -1, 32)
			strings[snum] = s
			snum++
		}
		fmt.Printf("%.2f,%.2f -> %d  -  %.2f\n ", f1, f2, flag, f)
	}
	fmt.Printf("预测成功:%d , 预测失败:%d \n", num1, num2)
	for _, y := range strings {
		if y == "" {
			continue
		}
		fmt.Println(y)
	}
}

func calc(x1, x2 float32) float32 {
	return relu(W11*x1+W21*x2+B1)*Q1 + relu(W12*x1+W22*x2+B2)*Q2 + relu(W13*x1+W23*x2+B3)*Q3
}

func relu(input float32) float32 {
	if input < 0 {
		return 0
	}
	return input
}
