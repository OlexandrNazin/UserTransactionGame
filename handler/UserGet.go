package handler

import (
	"UserTransactionGame/logger"
	"encoding/json"
	"net/http"
)

type UserGetStruct struct {
	UserId uint
	Token  string
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	var user UserGetStruct
	newDecoder := json.NewDecoder(r.Body)
	err := newDecoder.Decode(&user)
	if err != nil {
		logger.Logger(err)
		return
	}

	if user.Token == "" {

	}
	if user.UserId == 0 {

	}
}
