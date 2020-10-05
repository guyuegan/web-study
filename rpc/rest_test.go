package rpc

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"testing"
)

func TestRest(t *testing.T) {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	router.GET("/user/:id", getuser)
	router.POST("/user/:id", adduser)
	router.DELETE("/user/:id", deluser)
	router.PUT("/user/:id", moduser)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func moduser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	name := params.ByName("id")
	fmt.Fprintf(writer, "mod user "+name)
}

func deluser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	name := params.ByName("id")
	fmt.Fprintf(writer, "del user "+name)
}

func adduser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	name := params.ByName("id")
	fmt.Fprintf(writer, "add user "+name)
}

func getuser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	name := params.ByName("id")
	fmt.Fprintf(writer, "get user "+name)
}

func Hello(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	fmt.Fprintf(writer, "hello "+name)
}

func Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "welcome")
}
