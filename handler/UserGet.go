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
		logger.ResWriter(w, "error: fail")
		logger.Logger(err)
		return
	}

	if user.Token == "" {
		logger.ResWriter(w, "error: failToken")
	}
	if user.UserId == 0 {
		logger.ResWriter(w, "error: failUserId")
	}
}
