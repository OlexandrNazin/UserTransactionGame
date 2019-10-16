package handler

import (
	"UserTransactionGame/logger"
	"UserTransactionGame/memory"
	"encoding/json"
	"log"
	"net/http"
)

type UserGetStruct struct {
	Id    uint
	Token string
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	var user UserGetStruct
	newDecoder := json.NewDecoder(r.Body)
	err := newDecoder.Decode(&user)
	if err != nil {
		logger.ResWriter(w, map[string]interface{}{"error": "fail"})
		logger.Logger(err)
		return
	}
	log.Println(user)

	if user.Token == "" {
		logger.ResWriter(w, map[string]interface{}{"error": "failToken"})
	}
	if user.Id == 0 {
		logger.ResWriter(w, map[string]interface{}{"error": "failUserId"})
		return
	}
	info, err := GetUser(&user)
	if err != nil {
		logger.Logger(err)
	}
	logger.ResWriter(w, info)

}

func GetUser(User *UserGetStruct) (map[string]interface{}, error) {
	user, err := memory.Memory.GetUser(User.Id)
	if err != nil {
		return nil, err
	}
	userGlobalInfo := memory.Memory.GetUserInfo(user.Id)

	return map[string]interface{}{
		"id":           user.Id,
		"balance":      user.Balance,
		"depositCount": userGlobalInfo.DepositCount,
		"depositSum":   userGlobalInfo.DepositSum,
		"betCount":     userGlobalInfo.BetCount,
		"betSum":       userGlobalInfo.BetSum,
		"winCount":     userGlobalInfo.WinCount,
		"winSum":       userGlobalInfo.WinSum,
	}, nil
}
