package main

import (
	"./orange"
	"fmt"
)

func main() {
	// implicit assignment of unexported field 'firstName' in orange.Person literal
	//i := orange.Person{"11", "chen"}
	person := new(orange.Person)
	person.SetFirstName("hhaha")
	fmt.Println(person.FirstName())
}
