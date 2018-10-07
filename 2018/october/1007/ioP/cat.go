package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(line int, r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%d %s", line, buf)
		line++
	}

	return
}

func main() {
	line := 1
	flag.Parse()
	if flag.NArg() == 0 {
		cat(line, bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(line, bufio.NewReader(f))
	}
}
