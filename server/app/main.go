package main

import (
	"net/http"
	"todo/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Hello)
	http.ListenAndServe(":8000", r)
}
