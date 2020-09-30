package file

import (
	"fmt"
	"os"
	"testing"
)

const fileName = "1.txt"

func TestWrite(t *testing.T) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(fileName, err)
		return
	}
	defer file.Close()

	for i := 0; i < 10; i++ {
		file.WriteString("this is a test\r\n")
		file.Write([]byte("this is a test\r\n"))
	}
}

func TestRead(t *testing.T) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fileName, err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

/*
删除文件
Go 语言里面删除文件和删除文件夹是同一个函数
func Remove(name string) Error
调用该函数就可以删除文件名为 name 的文件
 */
func TestRemove(t *testing.T) {

}