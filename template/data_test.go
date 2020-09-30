package template

import (
	"html/template"
	"os"
	"testing"
)

type Person struct {
	Username string
	email string
}

/*
func handler(w http.ResponseWriter, r *http.Request) {
    t := template.New("some template") // 创建一个模板
    t, _ = t.ParseFiles("tmpl/welcome.html")  // 解析模板文件
    user := GetUser() // 获取当前用户信息
    t.Execute(w, user)  // 执行模板的 merger 操作
}
 */
func TestTpl(t *testing.T) {
	// 创建模板
	tpl := template.New("fieldname example")
	// 解析模板
	tpl, _ = tpl.Parse("hello {{.Username}} {{.email}}")
	// 填充数据(导出字段才有效)
	tpl.Execute(os.Stdout, Person{Username: "gopher", email: "xxx@yyy"})
}


type Friend struct {
	Fname string
}

type Person2 struct {
	Username string
	Emails []string
	Friends []*Friend
}

/*
嵌套对象：{{with …}}…{{end}} 和 {{range …}}{{end}} 来进行数据的输出。
{{range}} 这个和 Go 语法里面的 range 类似，循环操作数据
{{with}} 操作是指当前对象的值，类似上下文的概念
*/
func TestTplNest(t *testing.T) {
	f1 := Friend{Fname: "go1"}
	f2 := Friend{Fname: "go2"}
	tpl := template.New("fieldname nest example")
	tpl, _ = tpl.Parse(`hello {{.Username}}
		{{range .Emails}}
			an email {{.}}
		{{end}}
		{{with .Friends}}
			{{range .}}
				my friend is {{.}}
			{{end}}
		{{end}}
	`)
	p := Person2{Username: "gopher", Emails: []string{"x@y", "w@z"}, Friends: []*Friend{&f1, &f2}}
	tpl.Execute(os.Stdout, p)
}


