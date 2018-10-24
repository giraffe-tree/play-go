package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		// 直接用golang,不能用相对路径
		t, _ := template.ParseFiles("C:/Users/Administrator/Desktop/2018/August/learn/play-go/2018/october/1011/web/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		//for key, value := range r.Form {
		//	fmt.Println(key, value)
		//}
		//if len(r.Form["username"][0]) == 0 {
		//	fmt.Println("username can not be null")
		//}
		//i, err := strconv.Atoi(r.Form.Get("password"))
		//if err != nil {
		//	fmt.Println(err)
		//	fmt.Println(i)
		//} else {
		//	fmt.Println(i)
		//}
		//
		//if m, err := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		//	fmt.Println(err)
		//	fmt.Println(m)
		//}
		s := template.HTMLEscapeString(r.Form.Get("username"))
		fmt.Println("remove html: ", s)

	}
}

func main() {
	http.HandleFunc("/", sayhelloName2)      //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
