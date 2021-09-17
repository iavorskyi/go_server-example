package main

import (
	"log"
	"net/http"
	"os"

	"github.com/unrolled/render"

	"github.com/gorilla/mux"
)

var (
	firstVar  = os.Getenv("FIRST_VAR")
	secondVar = os.Getenv("SECOND_VAR")
	ren       = render.New()
)

func main() {

	log.Fatal(http.ListenAndServe(":8000", Handlers()))
}

func Handlers() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/root", rootHandler).Methods("GET")

	return r
}

func init() {
}
