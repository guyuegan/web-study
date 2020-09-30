package regex

import (
	"regexp"
	"testing"
)

func TestIpMatch(t *testing.T) {
	println(isIp("0.0.0.0"))

	matchString, _ := regexp.MatchString("(\\d{1,2}\\.){3}(\\d{1,2})", "0.0.0.0")
	println(matchString)

	println(isNum("123"))
	println(isNum("123%"))
	println(isNum(" 12"))
}

func isIp(ip string) bool {
	// todo sth wrong
	m, _ := regexp.MatchString("((\\d{1,2} | 1\\d\\d | 2[0-4][0-9] | 25[0-5])\\.){3}(\\d{1,2} | 1\\d\\d | 2[0-4][0-9] | 25[0-5])", ip)
	return m
}

func isNum(num string) bool {
	m, _ := regexp.MatchString("^\\d+$", num)
	return m
}