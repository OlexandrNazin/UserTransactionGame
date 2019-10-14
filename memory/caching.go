package memory

import (
	"UserTransactionGame/handler"
	"sync"
)

type memCache struct {
	mu               sync.Mutex
	itemUsers        map[uint]*handler.UserCreateStruct
	itemTransactions map[uint]*handler.UserTransactionStruct
}

var memory memCache

func init() {
	memory = memCache{
		itemUsers:        make(map[uint]*handler.UserCreateStruct),
		itemTransactions: make(map[uint]*handler.UserTransactionStruct),
	}

}
