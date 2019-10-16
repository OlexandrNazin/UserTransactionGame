package handler

import (
	"UserTransactionGame/logger"
	"UserTransactionGame/memory"
	"UserTransactionGame/model"
	"encoding/json"
	"log"
	"net/http"
)

type UserCreateStruct struct {
	Id      uint
	Balance float64
	Token   string
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user UserCreateStruct
	newDecoder := json.NewDecoder(r.Body)
	err := newDecoder.Decode(&user)
	log.Println(user)
	if err != nil {
		logger.ResWriter(w, map[string]interface{}{"error": "fail"})
		logger.Logger(err)
		return
	}
	User, err := createUser(&user)
	if err != nil {
		logger.ResWriter(w, map[string]interface{}{"error": "invalid Id"})
		return
	}

	err = memory.Memory.AddUserCache(User)
	if err != nil {
		logger.Logger(err)
		logger.ResWriter(w, map[string]interface{}{"error": err.Error()})
		return
	}

	if user.Token == "" {
		logger.ResWriter(w, map[string]interface{}{"error": "Invalid token"})
	}

}
func createUser(user *UserCreateStruct) (*model.UserCreate, error) {

	if user.Id == 0 {
		return nil, logger.WriteErr("invalid Id")
	}

	return &model.UserCreate{Id: user.Id, Balance: user.Balance}, nil
}
