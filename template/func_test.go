package template

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"testing"
)

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	split := strings.Split(s, "@")
	if len(split) != 2 {
		return s
	}

	// replace the @ by " at "
	return split[0] + " at " + split[1]
}

func TestFunc(t *testing.T) {
	tpl := template.New("func example")
	tpl = tpl.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	tpl, _ = tpl.Parse(`
			{{range .Emails}}
				an email {{.|emailDeal}}
			{{end}}`)
	p := Person2{Username: "gopher", Emails: []string{"x@y", "w@z"}}
	tpl.Execute(os.Stdout, p)
}
