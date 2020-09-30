package db_test

import (
	"fmt"
	"github.com/astaxie/goredis"
	"testing"
)

func Test5(t *testing.T) {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"

	// 字符串操作
	client.Set("go", []byte("redis"))
	val, _ := client.Get("go")
	fmt.Println(string(val))
	client.Del("go")

//	list操作
	vals := []string{"a", "b", "c", "d", "e"}
	for _, v := range vals {
		client.Rpush("ls", []byte(v))
	}
	dbvals, _ := client.Lrange("ls", 0, 4)
	for i, v := range dbvals {
		println(i, ":", string(v))
	}
	client.Del("ls")
}


