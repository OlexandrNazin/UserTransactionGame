package handler

import (
	"UserTransactionGame/logger"
	"encoding/json"
	"net/http"
)

type UserTransactionStruct struct {
	UserId        uint
	TrancactionId uint
	Bet           string
	Amount        float64
	Token         string
}

func UserTransaction(w http.ResponseWriter, r *http.Request) {
	var user UserTransactionStruct
	newDecoder := json.NewDecoder(r.Body)
	err := newDecoder.Decode(&user)
	if err != nil {
		logger.ResWriter(w, map[string]interface{}{"error": "fail"})
		logger.Logger(err)
		return
	}
	if user.Token == "" {

	}
	if user.Amount == 0.0 {

	}
	if user.UserId == 0 {

	}
	if user.TrancactionId == 0 {

	}
	if user.Bet == "" {

	}
}
