package json

import (
	"encoding/json"
	"os"
	"testing"
)

type Srv struct {
	//ID不会导出都json
	ID int `json:"-"`

	//SrvName2的值会进行二次 JSON 编码
	SrvName string `json:"srvName"`
	SrvName2 string `json:"srvName2,string"`

	//如果 SrvIp 为空，则不输出到 JSON 串中
	SrvIp string `json:"srvIp,omitempty"`
}

func TestJsonStructTag(t *testing.T) {
	srv := Srv{3, `GO "1.0"`, `GO "1.0"`, ``}
	marshal, _ := json.Marshal(srv)
	os.Stdout.Write(marshal)
}
