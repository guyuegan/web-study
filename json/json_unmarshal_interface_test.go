package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
bool 代表 JSON booleans,
float64 代表 JSON numbers,
string 代表 JSON strings,
nil 代表 JSON null.
 */

func TestJsonUnmarshalInterface(t *testing.T) {
	jsonStr := `{"Name":"Wednesday",
				"Age":6,
				"fav":{"name":"basketball", "time":"3 year"},
				"Parents":["Gomez","Morticia"]}`
	bytes := []byte(jsonStr)
	var jsonObj interface{}
	if err := json.Unmarshal(bytes, &jsonObj); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", jsonObj)

	parseJson(jsonObj)
}

func parseJson(jsonObj interface{}) {
	//通过断言方式访问嵌套结构
	jsonMap := jsonObj.(map[string]interface{})
	for key, val := range jsonMap {
		switch valType := val.(type) {
		case string:
			fmt.Println(key, "is string", valType)
		case int:
			fmt.Println(key, "is int", valType)
		case float64:
			fmt.Println(key, "is float64", valType)
		case []interface{}:
			fmt.Println(key, "is an array", valType)
			for i, e := range valType {
				fmt.Println(i, e)
			}
		case map[string]interface{}:
			fmt.Println(key, "is an map", valType)
			parseJson(val)
		default:
			fmt.Println(key, "is of a type I don't know how to handle")
		}
	}
}
