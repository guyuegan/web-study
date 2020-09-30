package template

import (
	"html/template"
	"os"
	"testing"
)

func TestCondition(t *testing.T) {
	tpl := template.New("tpl test")
	tpl = template.Must(tpl.Parse("空 pipeline if demo: {{if ``}} 不会输出 {{end}}\n"))
	tpl.Execute(os.Stdout, nil)

	tpl2 := template.New("tpl test")
	template.Must(tpl2.Parse("不为空的 pipeline if demo: {{if `anything`}} 我会输出 {{end}}\n"))
	tpl2.Execute(os.Stdout, nil)

	tpl3 := template.New("tpl test")
	tpl3 = template.Must(tpl3.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分 {{end}}\n"))
	tpl3.Execute(os.Stdout, nil)

	type Inventory struct {
		Material string
		Count    int
	}
	sweaters := Inventory{"axe", 0}
	tmpl, _ := template.New("test").Parse(`
		{{$a := .Count}} {{$b := 17}} {{$c := 18}}	
		{{if eq .Count $b}} 相等 {{else}} 不相等 {{end}}`)
	tmpl.Execute(os.Stdout, sweaters)
}


