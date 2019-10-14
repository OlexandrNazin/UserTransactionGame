package handler

import (
	"UserTransactionGame/logger"
	"encoding/json"
	"net/http"
)

type UserCreateStruct struct {
	Id      uint
	Token   string
	Balance float64
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user UserCreateStruct
	newDecoder := json.NewDecoder(r.Body)
	err := newDecoder.Decode(&user)
	if err != nil {
		logger.Logger(err)
		return
	}

	if user.Token == "" {

	}
	if user.Id == 0 {

	}
	if user.Balance == 0.0 {

	}
}
