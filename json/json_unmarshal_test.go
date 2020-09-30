package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type Server struct {
	ServerName string
	ServerIp string
}

type ServerSlice struct {
	Servers []Server
}

func TestJsonUnmarshal(t *testing.T) {
	file, err := os.Open("servers.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	servers := ServerSlice{}
	if err := json.Unmarshal(content, &servers); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", servers)
}
