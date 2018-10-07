package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func main() {
	test2()
}

func test2() {

	cmd := exec.Command("echo", "%GO_ROOT%")
	//  cmd.Output() 是对 Run() 的简单封装,返回输出buf
	buf, _ := cmd.Output() // 错误处理

	print(string(buf))
}

func test1() {
	command := exec.Command("java", "-version")
	//err := command.Run()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//} else {
	//	fmt.Println("success...")
	//}
	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("out:", out.String())

}

func TestAb(t *testing.T) {
	fmt.Println("hello", t)
}
