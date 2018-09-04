package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // 画布大小
	cells         = 100                 // 网格单元的个数
	xyrange       = 30.0                // 坐标轴范围
	xyscale       = width / 2 / xyrange // x,y轴上 每个单位长度的像素
	zscale        = height * 0.4        // z轴上每个单位长度的像素
	angle         = math.Pi / 6         // x,y轴的角度
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey;fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {

		}
	}

}
