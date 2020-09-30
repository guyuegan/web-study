package url_test

import (
	"fmt"
	"net/url"
	"testing"
)

func Test(t *testing.T) {
	v := url.Values{}
	v.Set("name", "ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	fmt.Println(v.Encode())
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
}

