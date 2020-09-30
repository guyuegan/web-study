package xml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Srv [] Server `xml:"server"`
	AuthorName string `xml:"author>name"`
	AuthorPhone string `xml:"author>phone"`
	OtherField string `xml:",any"`
	Description string `xml:",innerxml"`
}

type Server struct {
	XMLName xml.Name `xml:"server"`
	Name string `xml:"serverName"`
	Ip string `xml:"serverIP"`
}

func TestXmlUnmarshal(t *testing.T) {
	file, err := os.Open("servers.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	servers := &Servers{}
	if err := xml.Unmarshal(content, servers); err != nil {
		panic(err)
	}
	//fmt.Println(servers)
	fmt.Printf("%+v\n", servers)
}