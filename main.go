package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello web") //这个写入到w的是输出到客户端的
}

/**
Request 本身也提供了 FormValue () 函数来获取用户提交的参数。如 r.Form ["username"] 也可写成
r.FormValue ("username")。调用 r.FormValue 时会自动调用 r.ParseForm，所以不必提前调用。
r.FormValue 只会返回同名参数中的第一个，若参数不存在则返回空字符串。
*/
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("parseForm: ", err)
		}
		//	请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

//【坑】
// curl http://localhost:9090/ （无参调用，正常返回）
// curl http://localhost:9090/?url_long=111 （带参调用，no matches found）
// curl "http://localhost:9090/?url_long=111" （加上""，正常返回）
func main() {
	//要编写一个 Web 服务器很简单，只要调用 http 包的两个函数就可以了

	http.HandleFunc("/", SayHello)           //设置访问路由
	http.HandleFunc("/login", login)         //设置访问路由
	err := http.ListenAndServe(":9090", nil) //设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
