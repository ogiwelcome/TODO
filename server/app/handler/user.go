package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/domain"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	user := domain.NewUser("hoge", "fuga")
	respBody, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	_, err = w.Write(respBody)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	return
}
