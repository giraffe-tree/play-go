package main

import (
	"fmt"
	"time"
)

type Employee struct {
	id         int
	name       string
	address    string
	updateTime time.Time
	salary     int
	managerId  int
}

func main() {
	var dilbert Employee
	dilbert.salary = 18000
	//dilbert.address == nil?"null":dilbert.address

	fmt.Println(dilbert.address == "")
	fmt.Println(dilbert.salary)
	name := &dilbert.name
	*name = "Senior " + *name
	fmt.Println(dilbert.name)

}
