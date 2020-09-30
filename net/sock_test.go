package net

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func TestClient(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:7777")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	time.Sleep(time.Second)
	_, err = conn.Write([]byte("hi server"))
	checkError(err)
	time.Sleep(time.Second)
	echoMsg, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println("echo: ", string(echoMsg))
	os.Exit(0)
}

func TestServer(t *testing.T) {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	buf := make([]byte, 2048)
	for {
		println("listening ...")
		conn, err := listener.Accept()
		println("accept conn ...")
		if err != nil {
			continue
		}

		readLen, err := conn.Read(buf)
		checkError(err)
		// 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
		conn.Write([]byte(fmt.Sprintf("%s, now is %s",
			string(buf[:readLen]),
			time.Now().Format("2006-01-02 15:04:05"))))
		conn.Close()
	}
}

