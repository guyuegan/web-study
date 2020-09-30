package xml

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

type ServersXml struct {
	XMLName xml.Name   `xml:"servers"`
	Version string     `xml:"version,attr"`
	Srv []ServerXml    `xml:"server"`
	AuthorName string  `xml:"author>name"`
	AuthorPhone string `xml:"author>phone"`
}

type ServerXml struct {
	XMLName xml.Name `xml:"server"`
	Name string `xml:"serverName"`
	Ip string `xml:"serverIP"`
}

func TestXmlUnmarshalBetter(t *testing.T) {
	serversXml := &ServersXml{Version: "1", AuthorName: "gyg", AuthorPhone: "123456"}
	serversXml.Srv = append(serversXml.Srv, ServerXml{Name: "Shanghai_VPN", Ip: "127.0.0.1"})
	serversXml.Srv = append(serversXml.Srv, ServerXml{Name: "Beijing_VPN", Ip: "127.0.0.2"})

	xmlContent, err := xml.MarshalIndent(serversXml, "xml: ", "	")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v", string(xmlContent))
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(xmlContent)

	ioutil.WriteFile("serversXml.xml", xmlContent, 0644)
}