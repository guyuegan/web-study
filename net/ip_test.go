package net

import (
	"net"
	"testing"
)

func TestParseIp(t *testing.T) {
	ip := net.ParseIP("127.0.0.1")
	println(ip.String())
}
