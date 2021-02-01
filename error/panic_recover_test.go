package error

import "testing"

func TestPanicRecover(t *testing.T) {
	println(GetUser(9))
}

var User = map[int]string{
	1: "bajie",
	2: "wukong",
	3: "shasheng",
	4: "tangsheng",
}

func GetUser(uid int) (username string) {
	defer func() {
		if x := recover(); x != nil {
			username = "Not Found"
		}
	}()

	username = User[uid]
	return
}
