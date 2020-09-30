package file

import (
	"fmt"
	"os"
	"testing"
)

func TestDir(t *testing.T) {
	os.Mkdir("a", 0777)
	os.MkdirAll("a/b/c", 0777)

	if err := os.Remove("a"); err != nil {
		// remove a: directory not empty
		fmt.Println(err)
	}
	os.RemoveAll("a")
}

