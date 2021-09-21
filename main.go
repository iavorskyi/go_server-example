package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/unrolled/render"

	"github.com/gorilla/mux"
)

var (
	firstVar  = "1"
	secondVar = "2"
	ren       = render.New()
)

func main() {
	log.Println("Listening port: 8000!!!")
	log.Fatal(http.ListenAndServe(":8000", Handlers()))
	fmt.Println(hello())
}

func Handlers() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/root", rootHandler).Methods("GET")

	return r
}

func hello() string {
	return "Hlo Golang"
}
