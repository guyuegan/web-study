package net

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"
)

func send(conn net.Conn) {
	_, err := conn.Write([]byte("timestamp"))
	checkError(err)
	buf := make([]byte, 2048)
	readLen, err := conn.Read(buf)
	checkError(err)
	fmt.Println(string(buf[:readLen]))
}

func TestSend(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:1200")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	send(conn)
}

/*
使用goru实现连接并发处理
 */
func TestRecv(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1200")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		logs.Info("accept conn: ", conn)
		if err != nil {
			continue
		}
		go HandleClient(conn)
	}
}

func HandleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // 设置超时
	request := make([]byte, 2048)                         // set maxium request length to 128B to prevent flood attack
	defer conn.Close()

	logs.Info("handle client: ", conn)
	for {
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		logs.Info("read conn: ", conn)
		if readLen == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:readLen])) == "timestamp" {
			conn.Write([]byte(fmt.Sprintf("%s, now is %s",
				string(request[:readLen]),
				strconv.FormatInt(time.Now().Unix(), 10))))
		} else {
			// 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
			conn.Write([]byte(fmt.Sprintf("%s, now is %s",
				string(request[:readLen]),
				time.Now().Format("2006-01-02 15:04:05"))))
		}
		request = make([]byte, 128) // clear last read content
	}
}
