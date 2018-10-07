package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	path  = "D:/2018/august/go/play-go/2018/october/1007/ioP/hello.txt"
	path2 = "D:/2018/august/go/play-go/2018/october/1007/ioP/hello2.txt"
	path3 = "D:/2018/august/go/play-go/2018/october/1007/ioP/hello3.txt"
)

func main() {
	//bytes, e := ReadAll(path)
	//if e == nil {
	//	fmt.Println(string(bytes))
	//	fmt.Println("end")
	//} else {
	//	fmt.Println(e)
	//}
	//fmt.Println(getCurrentDirectory())
	//ReadBuf()
	//ReadGoFile(path)
	writeFile2()
}

func writeFile2(){
	outputFile, outputError := os.OpenFile(path3, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i:=0; i<10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}

func writeFile() {
	inputFile := path
	outputFile := path2
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	// 0777 为 unix 中的权限
	err = ioutil.WriteFile(outputFile, buf, 0777) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
func ReadBuf() {
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir //strings.Replace(dir, "\\", "/", -1)}
}

func ReadGoFile(path string) {
	file, e := os.Open(path)
	if e != nil {
		fmt.Println("open file ... error - ", e)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		fmt.Print(readString)
		if err == io.EOF {
			return
		}
	}
}
