package handler

import (
	"UserTransactionGame/logger"
	"UserTransactionGame/memory"
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
	CreatedTime   time.Time
	Token         string
}

func UserDeposit(w http.ResponseWriter, r *http.Request) {
	var userDeposit UserDepositStruct
	newDecoder := json.NewDecoder(r.Body)
	err := newDecoder.Decode(&userDeposit)
	if err != nil {
		logger.ResWriter(w, map[string]interface{}{"error": "fail"})
		logger.Logger(err)
		return
	}
	if userDeposit.Token == "" {
		logger.ResWriter(w, map[string]interface{}{"error": "Invalid token"})
	}
	if userDeposit.Amount == 0.0 {
		logger.ResWriter(w, map[string]interface{}{"error": "Invalid amount"})
	}
	if userDeposit.UserId == 0 {

	}
	if userDeposit.DepositId == 0 {
		logger.ResWriter(w, map[string]interface{}{"error": "Invalid deposit"})
	}
	deposit, err := memory.Memory.GetUser(userDeposit.UserId)
	if err != nil {
		logger.ResWriter(w, map[string]interface{}{"error": err.Error()})
		return
	}
	userDeposit.CreatedTime = time.Now()
	userDeposit.BalanceBefore = deposit.Balance
	userDeposit.BalanceAfter = deposit.Balance + userDeposit.Amount
	deposit.Balance += userDeposit.Amount

}
