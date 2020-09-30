package xml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type Class struct {
	Name string
	Stu []Stu
}

type Stu struct {
	Name string
	No int
}

func TestXmlMarshal(t *testing.T) {
	cls := Class{"c1", []Stu{{"小明", 1}, Stu{"小红", 2}}}
	xmlContent, err := xml.MarshalIndent(cls, "sm: ", "	")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(xmlContent))
}
