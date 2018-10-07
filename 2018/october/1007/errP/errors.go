// errors.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var err error

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s infile.txt outfile.txt", filepath.Base(os.Args[0]))
		fmt.Println(err)
		return
	}

	fmt.Printf("error: %" +
		"v", err)
	//switch err := errNotFound.(type) {
	//case os.PathError:
	//	fmt.Println(err)
	//}
	//
	
}
// error: Not found error