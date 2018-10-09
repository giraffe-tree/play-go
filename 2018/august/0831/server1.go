package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("192.168.2.127:8888", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	var url = " http://www.protontek.com/alipay/reportDetail.html?object_key=release/ios/2018/08/09/3787/15337941186.json&apiuid=3787&apitoken=3C82B887BACA4FE7E361EE2E75DA4074"

	var text = "<html>\n<body>\n<h1>我的第一个标题</h1>\n<a href=\"" + url + "\">Link text</a>\n</body>\n</html>"
	fmt.Fprintf(w, text)

}
