package template

import (
	"fmt"
	"html/template"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	s1, _ := template.ParseFiles("tmpl/header.tmpl", "tmpl/content.tmpl", "tmpl/footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)

}
