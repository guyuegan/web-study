package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func TestJsonRpcSrv(t *testing.T) {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

func TestJsonRpcCli(t *testing.T) {
	cli, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	//	sync call
	args := Args{17, 8}
	var reply int
	err = cli.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Arith.Multiply err: ", err)
	}
	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)

	var quo Quotient
	err = cli.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal("Arith.Divide err: ", err)
	}
	fmt.Printf("%d / %d = %d; %[1]d %% %[2]d = %[4]d", args.A, args.B, quo.Quo, quo.Rem)
}
