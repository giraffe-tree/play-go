package main

func main() {

}

func test1() {
	ch1 := make(chan string, 1000)
	<-ch1

}
