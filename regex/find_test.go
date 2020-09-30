package regex

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Test_find(t *testing.T) {
	a := "I am learning Go language"
	reg, _ := regexp.Compile("[a-z]{2,4}")

	// 查找符合正则的第一个
	one := reg.FindString(a)
	fmt.Println("FindString:", string(one))

	// 查找符合正则的所有 slice, n 小于 0 表示返回全部符合的字符串，不然就是返回指定的长度
	all := reg.FindAllString(a, -1)
	fmt.Println("FindAllString:", all)

	// 查找符合条件的 index 位置, 开始位置和结束位置
	index := reg.FindStringIndex(a)
	fmt.Println("FindStringIndex:", index)

	// 查找符合条件的所有的 index 位置，n 同上
	allIndex := reg.FindAllStringIndex(a, -1)
	fmt.Println("FindAllStringIndex:", allIndex)

	reg2, _ := regexp.Compile("am(.*)lang(.*)")

	// 查找 Submatch, 返回数组，第一个元素是匹配的全部元素，第二个元素是第一个 () 里面的，第三个是第二个 () 里面的
	// 下面的输出第一个元素是 "am learning Go language"
	// 第二个元素是 " learning Go "，注意包含空格的输出
	// 第三个元素是 "uage"

	submatch := reg2.FindStringSubmatch(a)
	fmt.Println(strings.Join(submatch, ","))

	// 定义和上面的 FindIndex 一样
	submatchIndex := reg2.FindStringSubmatchIndex(a)
	fmt.Println(submatchIndex)

	// FindAllSubmatch, 查找所有符合条件的子匹配
	allSubmatch := reg2.FindAllStringSubmatch(a, -1)
	for _, ele := range allSubmatch {
		fmt.Println(strings.Join(ele, ","))
	}

	// FindAllSubmatchIndex, 查找所有子匹配的 index
	allSubmatchIndex := reg2.FindAllStringSubmatchIndex(a, -1)
	fmt.Println(allSubmatchIndex)
}
