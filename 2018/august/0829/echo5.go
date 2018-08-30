package main

import (
	"os"
	"fmt"
)

func main() {

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i)
		fmt.Println(os.Args[i])
	}

}
