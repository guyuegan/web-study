package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/* 通过 struct tag 定义来实现输出小写首字母

type Server struct {
    ServerName string `json:"serverName"`
    ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
    Servers []Server `json:"servers"`
}
 */

func TestJsonMarshal(t *testing.T) {
	var serverSlice ServerSlice
	serverSlice.Servers = append(serverSlice.Servers, Server{"Shanghai_VPN", "127.0.0.1"})
	serverSlice.Servers = append(serverSlice.Servers, Server{"Beijing_VPN", "127.0.0.1"})
	marshal, err := json.Marshal(serverSlice)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(marshal))
}