package main

import (
	"net/http"
	"todo/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Hello)
	router.Use(handler.CORSOriginMiddleware)
	http.ListenAndServe(":8000", router)
}
