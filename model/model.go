package model

import "time"

type UserCreate struct {
	Id          uint
	Balance     float64
	CreatedTime time.Time
}
type UserTransaction struct {
	Id            uint
	UserId        uint
	BetId         uint
	Amount        float64
	BalanceBefore float64
	BalanceAfter  float64
	CreatedTime   time.Time
}
type UserDeposit struct {
	Id            int
	UserId        uint
	Amount        float64
	Token         string
	BalanceBefore float64
	BalanceAfter  float64
	CreatedTime   time.Time
}

type BetType string
type Bet struct {
	Id          uint
	Type        BetType
	UserId      uint
	Amount      float64
	CreatedTime time.Time
}
type GlobalInfo struct {
	DepositCount uint
	DepositSum   float64
	BetCount     uint
	BetSum       float64
	WinCount     uint
	WinSum       float64
}
