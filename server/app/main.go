package main

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":3000", nil)
}
