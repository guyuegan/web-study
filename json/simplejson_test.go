package json

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"testing"
)


func TestSimpleJson(t *testing.T) {
	jsonStr, _ := simplejson.NewJson([]byte(`{
    "test": {
        "array": [1, "2", 3],
        "int": 10,
        "float": 5.150,
        "bignum": 9223372036854775807,
        "string": "simplejson",
        "bool": true
    }
	}`))

	arr, _ := jsonStr.Get("test").Get("array").Array()
	i, _ := jsonStr.Get("test").Get("int").Int()
	str := jsonStr.Get("test").Get("string").MustString()
	fmt.Println(arr, i, str)
}
