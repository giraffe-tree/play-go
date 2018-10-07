package main

import (
	"fmt"
	"reflect"
)

type Bed struct {
	Size   int
	height int
	weight int
	length int
}

func (b Bed) Sum() {
	fmt.Println(b.weight * b.height * b.length)
}

func main() {
	bed := Bed{1, 11, 12, 13}
	value := reflect.ValueOf(bed)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
	/**
	方法首字母不大写无法调用
	 */
	fmt.Println("method num:", value.NumMethod())
	for j := 0; j < value.NumMethod(); j++ {
		value.Method(j).Call(nil)
	}

	e := reflect.ValueOf(&bed).Elem()
	if e.CanSet() {
		e.Field(0).SetInt(100)
		//reflect.Value.SetInt using value obtained using unexported field
		// field 首字母大写才可以设置
		//e.Field(1).SetInt(101)
		field1:= e.Field(1)
		if !field1.CanSet() {
			fmt.Println("field1 首字母大写才可以设置value")
		}
	}
	fmt.Println(bed.Size)
}
