package handler

import (
	"UserTransactionGame/logger"
	"UserTransactionGame/memory"
	"encoding/json"
	"net/http"
)

type UserTransactionStruct struct {
	UserId        uint
	TransactionId uint
	Type          string
	Amount        float64
	Token         string
}

func UserTransaction(w http.ResponseWriter, r *http.Request) {
	var Transaction UserTransactionStruct
	newDecoder := json.NewDecoder(r.Body)
	err := newDecoder.Decode(&Transaction)
	if err != nil {
		logger.ResWriter(w, map[string]interface{}{"error": "fail"})
		logger.Logger(err)
		return
	}
	if Transaction.Token == "" {
		logger.ResWriter(w, map[string]interface{}{"error": "invalid transaction token"})
	}
	if Transaction.Amount == 0.0 {
		logger.ResWriter(w, map[string]interface{}{"error": "invalid transaction amount"})
	}
	if Transaction.UserId == 0 {
		logger.ResWriter(w, map[string]interface{}{"error": "invalid transaction user id"})

	}
	if Transaction.TransactionId == 0 {
		logger.ResWriter(w, map[string]interface{}{"error": "invalid transaction id"})
	}
	if Transaction.Type != "Bet" && Transaction.Type != "Win" {
		logger.ResWriter(w, map[string]interface{}{"error": "invalid type transaction"})
		return
	}
	user, err := memory.Memory.GetUser(Transaction.UserId)
	if err != nil {
		logger.ResWriter(w, map[string]interface{}{"error": err.Error()})
		return
	}
	if Transaction.Type == "Bet" && Transaction.Amount > user.Balance {
		logger.ResWriter(w, map[string]interface{}{"error": "not enough funds on the balance sheet"})
		return
	}
	if Transaction.Type == "Bet" {
		user.Balance -= Transaction.Amount
		memory.Memory.UpdateTransactBet(user.Id, Transaction.Amount)
	}
	if Transaction.Type == "Win" {
		user.Balance += Transaction.Amount
		memory.Memory.UpdateTransactWin(user.Id, Transaction.Amount)
	}
	memory.Memory.ModUser(user)
	logger.ResWriter(w, map[string]interface{}{"error": "", "balance": user.Balance})
}
