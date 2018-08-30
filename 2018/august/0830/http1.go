package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	var s = "http://localhost:9000/user/register.mobile.captcha/1.0.0/get?mobile="
	var mobile = "xxxxx"
	resp, err :=http.Get(s + mobile)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))

}
