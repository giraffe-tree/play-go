package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(5e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	fmt.Println("add Washington")
	ch <- "Tripoli"
	fmt.Println("add Tripoli")
	ch <- "London"
	fmt.Println("add London")
	ch <- "Beijing"
	fmt.Println("add Beijing")
	ch <- "Tokyo"
	fmt.Println("add Tokyo")
}

func getData(ch chan string) {
	var input string
	//time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("get %s \n", input)
		time.Sleep(5e8)
	}
}
