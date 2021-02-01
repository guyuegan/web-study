package error

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"html/template"
	"net/http"
	"strings"
	"testing"
)

func serveHttp(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello":
		hello(w, r)
	case "/err404":
		notFound404(w, r)
	case "/err503":
		systemErr503(w, r)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
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

func notFound404(w http.ResponseWriter, r *http.Request) {
	logs.Error("页面找不到")
	tpl, _ := template.ParseFiles("404.html")
	ErrorInfo := "文件找不到"
	tpl.Execute(w, ErrorInfo)
}

func systemErr503(w http.ResponseWriter, r *http.Request) {
	logs.Critical("系统错误")
	tpl, _ := template.ParseFiles("503.html")
	ErrorInfo := "系统暂时不可用"
	tpl.Execute(w, ErrorInfo)
}

func TestErr404(t *testing.T) {
	http.HandleFunc("/", serveHttp)
	http.ListenAndServe(":9999", nil)
}
