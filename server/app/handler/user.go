package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/domain"
	"todo/usecase"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	user := domain.NewUser("hoge", "fuga")
	respBody, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
		fmt.Println(err)
	}

	_, err = w.Write(respBody)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := domain.NewUser("81f981b2-bdfa-4b98-b1a3-b4669f948a111", "鈴木ジオ")
	result, err := usecase.AddUser(&user)
	if err != nil {
		fmt.Println(err)
	}
	return
}
