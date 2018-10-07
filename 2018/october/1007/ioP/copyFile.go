package main

import (
	"fmt"
	"io"
	"os"
)

const (
	path4 = "D:/2018/august/go/play-go/2018/october/1007/ioP/hello.txt"
	path5 = "D:/2018/august/go/play-go/2018/october/1007/ioP/hello2.txt"
)

func main() {
	CopyFile(path5, path4)
	fmt.Println("Copy done!")
}
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
