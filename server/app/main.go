package main

import (
	"net/http"
	"os"
	"todo/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Hello)
	router.Use(handler.CORSOriginMiddleware)
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
