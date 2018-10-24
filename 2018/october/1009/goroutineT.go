package main

import (
	"fmt"
	"runtime"
)

/**

获取 goroutine 的状态

*/
func main() {

	runtime.GOMAXPROCS(4)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

}
