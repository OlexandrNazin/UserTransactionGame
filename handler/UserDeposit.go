package handler

import (
	"UserTransactionGame/logger"
	"encoding/json"
	"net/http"
	"time"
)

type UserDepositStruct struct {
	DepositId     int
	UserId        uint
	Amount        float64
	BalanceBefore float64
	BalanceAfter  float64
	CreatedAt     time.Time
	Token         string
}

func UserDeposit(w http.ResponseWriter, r *http.Request) {
	var user UserDepositStruct
	newDecoder := json.NewDecoder(r.Body)
	err := newDecoder.Decode(&user)
	if err != nil {
		logger.ResWriter(w, "error: fail")
		logger.Logger(err)
		return
	}
	if user.Token == "" {

	}
	if user.Amount == 0.0 {

	}
	if user.UserId == 0 {

	}
	if user.DepositId == 0 {

	}
}
