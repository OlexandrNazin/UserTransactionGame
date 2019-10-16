package memory

import (
	"UserTransactionGame/logger"
	"UserTransactionGame/model"
	"log"
	"sync"
)

type memCache struct {
	mu               sync.Mutex
	itemUsers        map[uint]*model.UserCreate
	itemTransactions map[uint]*model.UserTransaction
	itemGlobalInfo   map[uint]*model.GlobalInfo
	itemModified     map[uint]*model.UserCreate
}

var Memory memCache

func init() {
	Memory = memCache{
		itemUsers:        make(map[uint]*model.UserCreate),
		itemTransactions: make(map[uint]*model.UserTransaction),
		itemGlobalInfo:   make(map[uint]*model.GlobalInfo),
		itemModified:     make(map[uint]*model.UserCreate),
	}

}

func (mem *memCache) AddUserCache(NewUser *model.UserCreate) error {
	Memory.mu.Lock()
	defer mem.mu.Unlock()
	_, ok := mem.itemUsers[NewUser.Id]
	if ok {
		return logger.WriteErr("Cannot create user")
	}
	mem.itemUsers[NewUser.Id] = NewUser
	log.Println("New User")
	return nil
}
func (mem *memCache) GetUser(Id uint) (*model.UserCreate, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	user, ok := mem.itemUsers[Id]
	if ok {
		return user, nil
	}
	return nil, logger.WriteErr("User does not exist")
}

func (mem *memCache) GetUserInfo(id uint) *model.GlobalInfo {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	log.Println(id)
	_, err := mem.itemGlobalInfo[id]
	if err != true {
		mem.itemGlobalInfo[id] = &model.GlobalInfo{
			DepositCount: 0,
			DepositSum:   0,
			BetCount:     0,
			BetSum:       0,
			WinCount:     0,
			WinSum:       0,
		}
	}
	return mem.itemGlobalInfo[id]
}

func (mem *memCache) UpdateDeposit(id uint, sum float64) {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	_, err := mem.itemGlobalInfo[id]
	if !err {
		mem.itemGlobalInfo[id] = &model.GlobalInfo{}
	}
	info := mem.itemGlobalInfo[id]
	info.DepositCount++
	info.DepositSum += sum
}

func (mem memCache) UpdateTransactBet(id uint, sum float64) {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	_, err := mem.itemGlobalInfo[id]
	if !err {
		mem.itemGlobalInfo[id] = &model.GlobalInfo{}
	}
	info := mem.itemGlobalInfo[id]
	info.DepositCount++
	info.DepositSum += sum

}
func (mem memCache) UpdateTransactWin(id uint, sum float64) {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	_, err := mem.itemGlobalInfo[id]
	if !err {
		mem.itemGlobalInfo[id] = &model.GlobalInfo{}
	}
	info := mem.itemGlobalInfo[id]
	info.DepositCount++
	info.DepositSum += sum
}

func (mem *memCache) ModUser(user *model.UserCreate) {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	mem.itemModified[user.Id] = user
}
