package string

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
//	Contains
	println(strings.Contains("seafood", "foo"))
	println(strings.Contains("seafood", strings.ToLower("Fo")))
	println(strings.Contains("seafood", "Fo"))

	println(strings.Contains("seafood", "bar"))
	println(strings.Contains("seafood", ""))
	println(strings.Contains("", ""))

//	Join
	println(strings.Join([]string{"foo", "bar", "baz"}, "#"))

//	Split
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "ok, let's go"))

	//	Index
	println(strings.Index("chicken", "ken"))
	println(strings.Index("chicken", "dmr"))

//	Repeat
	println("ba" + strings.Repeat("na", 2))

//	Replace
	println(strings.Replace("oink oink oink", "k", "ky", 2))
	println(strings.Replace("oink oink oink", "oink", "moo", -1))

//	Trim
	println(strings.Trim(" !!!$$$ attention !!!$$$ ", "!$ "))

//	Fields
	fmt.Printf("fields are %q", strings.Fields(" foo bar baz "))
}

func TestStrConv(t *testing.T) {
//	Append系列， make的第三个参数：https://blog.csdn.net/weiyuefei/article/details/77968699
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcd")
	str = strconv.AppendQuoteRune(str, '单')
	println(string(str))

//	Format系列
	println(strconv.FormatBool(false))
	println(strconv.FormatFloat(123.23, 'g', 12, 64))
	println(strconv.FormatInt(7, 2))
	println(strconv.FormatUint(16, 8))
	println(strconv.Itoa(12345))

//	Parse系列 [关键字也能做变量名哦，厉害了]
	bool, _ := strconv.ParseBool("false")
	println(bool)
	float, _ := strconv.ParseFloat("123.23", 64)
	println(float)
	int, _ := strconv.ParseInt("1234", 10, 64)
	println(int)
	uint, _ := strconv.ParseUint("12345", 10, 64)
	println(uint)
	atoi, _ := strconv.Atoi("1024")
	println(atoi)
}

func TestMake(t *testing.T) {
	a := make([]int, 5, 10)
	fmt.Printf("%d, %d, %v\n", len(a), cap(a), a)
	b := a[:cap(a)]
	fmt.Printf("%d, %d, %v\n", len(b), cap(b), b)
	b = append(b, 11)
	fmt.Printf("%d, %d, %v\n", len(b), cap(b), b)
}