package rpc

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"testing"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func TestHttpRpcSrv(t *testing.T) {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	if err := http.ListenAndServe(":1234", nil); err != nil {
		fmt.Println(err.Error())
	}
}

func TestHttpRpcCli(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dial: ", err)
	}

	//	sync call
	args := Args{17, 8}
	var reply int
	if err = client.Call("Arith.Multiply", args, &reply); err != nil {
		log.Fatal("call Arith.Multiply: ", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n\n", args.A, args.B, reply)

	var quo Quotient
	if err = client.Call("Arith.Divide", args, &quo); err != nil {
		log.Fatal("Arith.Divide: ", err)
	}
	fmt.Printf("Arith: %d / %d = %d, %[1]d %% %[2]d = %[4]d\n", args.A, args.B, quo.Quo, quo.Rem)
}
