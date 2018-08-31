package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get("http://" + url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)

		b, err := io.Copy(os.Stdout, resp.Body)
		s := resp.Status
		fmt.Println(s)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)

	}

}
