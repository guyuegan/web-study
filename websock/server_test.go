package websock

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"testing"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("can't receive")
			break
		}
		fmt.Println("received back from client: " + reply)

		msg := "received: " + reply
		fmt.Println("sending to client: " + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("can't send")
			break
		}
	}
}

func TestEchoSrv(t *testing.T) {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}

func TestSendMsg(t *testing.T) {
	http.Handle("/", websocket.Handler(Send))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}

func Send(ws *websocket.Conn) {
	var name string
	for {
		fmt.Scanln(&name)
		if err := websocket.Message.Send(ws, "hello ws, i'm "+name); err != nil {
			fmt.Println("can't send")
			return
		}
	}
}
