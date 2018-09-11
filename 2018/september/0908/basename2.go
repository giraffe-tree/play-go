package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(basename("a.go"))
	fmt.Println(basename("ss/ssfds/a.g.go"))


}

func basename(s string) string {

	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, ".");dot>=0 {
		s = s[:dot]
	}
	return s

}
