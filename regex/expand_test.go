package regex

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_expand(t *testing.T) {
	src := []byte(`
		call hello alice
		hello bob	
		call hello eve
	`)
	pattern := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	var res []byte
	for _, s := range pattern.FindAllSubmatchIndex(src, -1) {
		res = pattern.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))

	var res2 []byte
	for _, s2 := range pattern.FindAllSubmatchIndex(src, -1) {
		res2 = pattern.Expand(res2, []byte("$2('$3')\n"), src, s2)
	}
	fmt.Println(string(res2))
}
